package main

import (
	"fmt"

	"github.com/liliang-cn/pipeit"
)

func main() {
	// Create a new process manager for 'gemini'
	// Just running help to see if it captures output
	pm := pipe.New("gemini", "--help")
    pm.SetOutputHandler(func(data []byte) {
        fmt.Print(string(data))
    })

	// Start using pipes (non-interactive mode might be safer for help)
	if err := pm.StartWithPipes(); err != nil {
		fmt.Printf("Error starting with pipes: %v\n", err)
        // try PTY fallback
        if err := pm.StartWithPTY(); err != nil {
            panic(err)
        }
	} else {
        defer pm.Stop()
    }

	pm.Wait()
}
