package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/liliang-cn/pipeit"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <command> [args...]\n", os.Args[0])
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	// Create a new process manager
	pm := pipe.New(command, args...)

	// Handle output
	pm.SetOutputHandler(func(data []byte) {
		os.Stdout.Write(data)
	})

	// Handle error (if separated, though PTY merges them usually)
	pm.SetErrorHandler(func(data []byte) {
		os.Stderr.Write(data)
	})

	// Start the process with PTY
	if err := pm.StartWithPTY(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting command: %v\n", err)
		os.Exit(1)
	}
	defer pm.Stop()

	// Forward Stdin to the process
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if n > 0 {
				_, wErr := pm.Write(buf[:n])
				if wErr != nil {
					return
				}
			}
			if err != nil {
				if err != io.EOF {
					// fmt.Fprintf(os.Stderr, "Stdin read error: %v\n", err)
				}
				return
			}
		}
	}()

	// Handle signals (pass interrupt to child if possible, or just exit)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for sig := range sigChan {
			// ideally we send this to the child process
			// pipe package might need a Signal method, but for now we stop.
			if sig == syscall.SIGINT {
				// User hit Ctrl+C, usually PTY handles this if we were in raw mode,
				// but since we are not in raw mode, our shell catches it.
				// For now, let's just stop the process manager.
				pm.Stop()
				os.Exit(0)
			}
		}
	}()

	// Wait for the process to finish
	if err := pm.Wait(); err != nil {
		// Verify if it's just an exit code error
		if exitErr, ok := err.(*os.PathError); ok {
			fmt.Fprintf(os.Stderr, "Command failed: %v\n", exitErr)
		}
	}
}
