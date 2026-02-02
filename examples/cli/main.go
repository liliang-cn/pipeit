package main

import (
	"fmt"
	"os"
	"time"

	"pipeit"
)

func main() {
	fmt.Println("=== CLI IO Capture Tool ===")

	// Get the example to run from environment variables
	example := os.Getenv("EXAMPLE")

	switch example {
	case "bash":
		runBash()
	case "zsh":
		runZsh()
	case "python":
		runPython()
	case "stream":
		runStream()
	default:
		printUsage()
		runBash() // Default to running bash example
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  EXAMPLE=bash go run *.go       # Run bash example")
	fmt.Println("  EXAMPLE=zsh go run *.go        # Run zsh example")
	fmt.Println("  EXAMPLE=python go run *.go     # Run python example")
	fmt.Println("  EXAMPLE=stream go run *.go     # Stream processing example")
	fmt.Println()
}

// runBash runs bash and captures output
func runBash() {
	fmt.Println("--- Bash Example ---")

	pm := pipe.New("bash", "--norc")

	// Set output handler
	pm.SetOutputHandler(func(data []byte) {
		fmt.Print(string(data))
	})

	if err := pm.StartWithPTY(); err != nil {
		fmt.Printf("Start failed: %v\n", err)
		return
	}
	defer pm.Stop()

	// Send command
	pm.Writeln("echo 'Hello from bash!'")
	time.Sleep(300 * time.Millisecond)

	pm.Writeln("pwd")
	time.Sleep(300 * time.Millisecond)

	pm.Writeln("exit")
	pm.Wait()
}

// runZsh runs zsh and captures output
func runZsh() {
	fmt.Println("--- Zsh Example ---")

	pm := pipe.New("zsh", "--no-rcs")

	pm.SetOutputHandler(func(data []byte) {
		fmt.Print(string(data))
	})

	if err := pm.StartWithPTY(); err != nil {
		fmt.Printf("Start failed: %v\n", err)
		return
	}
	defer pm.Stop()

	pm.Writeln("echo 'Hello from zsh!'")
	time.Sleep(300 * time.Millisecond)

	pm.Writeln("which zsh")
	time.Sleep(300 * time.Millisecond)

	pm.Writeln("exit")
	pm.Wait()
}

// runPython runs Python and interacts
func runPython() {
	fmt.Println("--- Python Example ---")

	pm := pipe.New("python3", "-q")

	pm.SetOutputHandler(func(data []byte) {
		fmt.Print(string(data))
	})

	if err := pm.StartWithPTY(); err != nil {
		fmt.Printf("Start failed: %v\n", err)
		return
	}
	defer pm.Stop()

	pm.Writeln("print('Hello from Python!')")
	time.Sleep(300 * time.Millisecond)

	pm.Writeln("import sys")
	time.Sleep(200 * time.Millisecond)

	pm.Writeln("print(sys.version)")
	time.Sleep(300 * time.Millisecond)

	pm.Writeln("exit()")
	pm.Wait()
}

// runStream stream processing example
func runStream() {
	fmt.Println("--- Stream Processing Example ---")

	pm := pipe.New("bash", "--norc")

	// Collect all output
	var output string
	pm.SetOutputHandler(func(data []byte) {
		str := string(data)
		output += str
		fmt.Print(str)
	})

	if err := pm.StartWithPTY(); err != nil {
		fmt.Printf("Start failed: %v\n", err)
		return
	}
	defer pm.Stop()

	// Run loop
	pm.Writeln("for i in 1 2 3; do echo \"Count: $i\"; done")
	time.Sleep(500 * time.Millisecond)

	pm.Writeln("exit")
	pm.Wait()

	fmt.Printf("\nCollected full output length: %d bytes\n", len(output))
}
