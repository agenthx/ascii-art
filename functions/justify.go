package piscine

import (
	"fmt"
	"os"
)

func Justify(str, lines []string, noOfS, noSpaceLength []int) {
	// Create a slice to store the lines for the current row
	var lineSlice string
	var space string

	// Iterate over the lines
	for i, line := range lines {
		if line == "" {
			fmt.Println()
		} else {
			space = Ss(i, noOfS[i], noSpaceLength[i])
			// Iterate over the 8 rows for each line
			for i := 0; i < 8; i++ {
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
					} else {
						lineSlice += space
					}
				}
				fmt.Println("|" + lineSlice + "|")
				lineSlice = ""
			}
		}
	}
}

func Ss(i, noOfS, noSpaceLength int) string {
	var space string
	if noOfS == 0 {
		return ""
	}
	for j := 2; j < (width()-noSpaceLength)/noOfS; j++ {
		space += " "
	}
	return space
}
