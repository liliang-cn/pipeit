package main

import (
	"fmt"
	"os"
	"time"

	"github.com/liliang-cn/pipeit"
)

func main() {
	fmt.Println("Starting Claude via pipeit...")

	// Create a new process manager for 'claude' using Config to set Env
	config := pipe.Config{
		Command: "claude",
		Env:     append(os.Environ(), "TERM=xterm-256color"),
		OnOutput: func(data []byte) {
			fmt.Print(string(data))
		},
	}

	pm := pipe.NewWithConfig(config)

	// Start the process with a PTY for interactive behavior
	if err := pm.StartWithPTY(); err != nil {
		panic(err)
	}
	defer pm.Stop()

	// Set terminal size - CRITICAL for interactive menus
	pm.SetWindowSize(24, 80)

	// Wait for initialization
	time.Sleep(3 * time.Second)

	// Use KeyEnter for confirmation
	fmt.Println("\n[PIPEIT]: Confirming workspace trust...")
	pm.WriteString(pipe.KeyEnter)

	// Wait for actual startup
	time.Sleep(8 * time.Second)

	// Send a simple prompt
	fmt.Println("\n[PIPEIT]: Sending prompt...")
	pm.Writeln("Briefly tell me who you are.")
	
	// Wait longer for response generation
	fmt.Println("[PIPEIT]: Waiting for response (45s)...")
	time.Sleep(45 * time.Second)

	fmt.Println("\n[PIPEIT]: Stopping...")
}
