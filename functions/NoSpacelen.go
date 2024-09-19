package piscine

import (
	"fmt"
	"os"
	"strings"
)

func NoSpacelen(str, lines []string) (int, int) {
	// Create a slice to store the lines for the current row
	var lineSlice string
	
	fmt.Println(lines)
	var words string
	for _, ch := range lines {
		words += string(ch)
	}
	arrWords := strings.Split(words, " ")
	fmt.Println(len(arrWords) - 1)
	// Iterate over the lines
	for _, line := range lines {
		if line == "" {
			fmt.Println()
		} else {
			// Iterate over the 8 rows for each line
			for i := 0; i < 1; i++ {
				// Iterate over the characters in the current line
				for _, ch := range line {
					// Check if the character is a valid ASCII character
					if ch < 32 || ch > 128 {
						fmt.Println("error input")
						os.Exit(7)
					}
					// Calculate the line index for the current character
					lineIndex := 1 + int(ch-32)*9 + i
					// Append the corresponding string to the line slice
					if ch != ' ' {
						lineSlice += str[lineIndex]
					}
				}
			}
		}
	}
	return len(lineSlice), 
}
