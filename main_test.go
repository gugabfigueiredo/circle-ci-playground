package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// Create a pipe to capture stdout
	r, w, _ := os.Pipe()
	originalStdout := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = originalStdout
	}()

	// Run the main function
	main()

	// Close the write end of the pipe
	w.Close()

	// Read the captured output from the read end of the pipe
	capturedOutput, _ := ioutil.ReadAll(r)

	// Check the output
	expectedOutput := "Hello, World!\n"
	actualOutput := string(capturedOutput)
	if actualOutput != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, actualOutput)
	}
}
