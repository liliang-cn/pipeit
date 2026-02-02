// How to use the pipeit library in your Go project
//
// 1. Install dependency:
//    go get pipeit/pkg/pipe
//
// 2. Import in your code:
//    import "pipeit/pkg/pipe"

package main

import (
	"fmt"
	"time"

	"pipeit"
)

// Example 1: Basic Usage
func Example1_Basic() {
	// Create process manager
	pm := pipe.New("bash", "--norc")

	// Set output handler
	pm.SetOutputHandler(func(data []byte) {
		fmt.Printf("[Output]: %s", string(data))
	})

	// Start
	pm.StartWithPTY()
	defer pm.Stop()

	// Send command
	pm.Writeln("echo 'Hello World'")
	time.Sleep(300 * time.Millisecond)

	pm.Writeln("exit")
	pm.Wait()
}

// Example 2: Collect output to variable
func Example2_CollectOutput() {
	pm := pipe.New("bash", "--norc")

	var allOutput string
	pm.SetOutputHandler(func(data []byte) {
		allOutput += string(data)
	})

	pm.StartWithPTY()
	defer pm.Stop()

	pm.Writeln("echo 'This will be collected'")
	time.Sleep(300 * time.Millisecond)

	pm.Writeln("exit")
	pm.Wait()

	fmt.Println("Collected output:", allOutput)
}

// Example 3: Create with configuration
func Example3_WithConfig() {
	pm := pipe.NewWithConfig(pipe.Config{
		Command: "bash",
		Args:    []string{"--norc"},
		Env:     []string{"MY_VAR=value"},
		OnOutput: func(data []byte) {
			fmt.Printf("[Output]: %s", string(data))
		},
	})

	pm.StartWithPTY()
	defer pm.Stop()

	pm.Writeln("echo $MY_VAR")
	time.Sleep(300 * time.Millisecond)

	pm.Writeln("exit")
	pm.Wait()
}

// Example 4: Interact with Python
func Example4_Python() {
	pm := pipe.New("python3", "-q")

	pm.SetOutputHandler(func(data []byte) {
		fmt.Printf("[Python]: %s", string(data))
	})

	pm.StartWithPTY()
	defer pm.Stop()

	pm.Writeln("print('Hello from Python')")
	time.Sleep(200 * time.Millisecond)

	pm.Writeln("2 + 2")
	time.Sleep(200 * time.Millisecond)

	pm.Writeln("exit()")
	pm.Wait()
}

// Example 5: Check process status
func Example5_ProcessStatus() {
	pm := pipe.New("bash", "--norc")

	pm.StartWithPTY()
	defer pm.Stop()

	fmt.Printf("Process running: %v\n", pm.IsRunning())
	fmt.Printf("Process PID: %d\n", pm.Pid())

	pm.Writeln("exit")
	pm.Wait()

	fmt.Printf("Process running: %v\n", pm.IsRunning())
}

// Example 6: Stream processing
func Example6_StreamProcessing() {
	pm := pipe.New("bash", "--norc")

	pm.SetOutputHandler(func(data []byte) {
		// Process each line of output in real-time
		str := string(data)
		// Here you can parse, save, forward, etc.
		fmt.Printf("[Real-time]: %s", str)
	})

	pm.StartWithPTY()
	defer pm.Stop()

	pm.Writeln("for i in {1..3}; do echo \"Count: $i\"; sleep 0.1; done")
	time.Sleep(500 * time.Millisecond)

	pm.Writeln("exit")
	pm.Wait()
}

// Example 7: Error handling
func Example7_ErrorHandling() {
	pm := pipe.New("bash", "--norc")

	pm.SetOutputHandler(func(data []byte) {
		fmt.Printf("[STDOUT]: %s", string(data))
	})

	pm.SetErrorHandler(func(data []byte) {
		fmt.Printf("[STDERR]: %s", string(data))
	})

	pm.StartWithPTY()
	defer pm.Stop()

	pm.Writeln("echo 'Normal output'")
	time.Sleep(200 * time.Millisecond)

	pm.Writeln("echo 'Error output' >&2")
	time.Sleep(200 * time.Millisecond)

	pm.Writeln("exit")
	pm.Wait()
}

// Example 8: Writef formatted write
func Example8_Writef() {
	pm := pipe.New("bash", "--norc")

	pm.SetOutputHandler(func(data []byte) {
		fmt.Printf("[Output]: %s", string(data))
	})

	pm.StartWithPTY()
	defer pm.Stop()

	// Use formatted write
	pm.Writef("echo 'Number: %d'\n", 42)
	time.Sleep(200 * time.Millisecond)

	pm.Writeln("exit")
	pm.Wait()
}

func main() {
	// Run any example
	Example1_Basic()
}
