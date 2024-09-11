package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	num := returnTerminalCols
	fmt.Println(num)
}

func returnTerminalCols() int {
	// Create a new command to run the "tput cols" command.
	cmd := exec.Command("tput", "cols")

	// Set the standard input of the command to os.Stdin (standard input).
	cmd.Stdin = os.Stdin

	// Execute the command and capture its output.
	out, err := cmd.Output()
	// Check if there was an error executing the command.
	if err != nil {
		// If there was an error, log the error message and return 0 as the terminal width.
		log.Fatal(err)
		return 0
	}
	// Convert the output (a byte slice) to a string and remove the trailing newline character.
	// Then, convert the string to an integer (terminal width).
	val, err := strconv.Atoi(string(out[:len(out)-1]))
	// Check if there was an error during the conversion.
	if err != nil {
		// If there was an error, return 0 as the terminal width.
		return 0
	}

	// Return the terminal width as an integer.
	return val
}
