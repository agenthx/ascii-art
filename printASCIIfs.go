package piscine

import (
	"fmt"
	"regexp"
	"strings"
)

func PrintASCIIfs(str []string, input string) {
	//if input is empty return nothing
	if input == "" {
		return
	}
	// Replace all "\n" in the input string with an actual newline
	input2 := strings.ReplaceAll(input, "\\n", "\n")
	//check if the input consists of only \n
	reN := regexp.MustCompile(`[^\\n]`)
	if !reN.MatchString(input) {
		fmt.Print(input2)
		return
	}
	// Split the input string into lines
	lines := strings.Split(input2, "\n")
	// Iterate over the lines
	for _, line := range lines {
		if line == "" {
			fmt.Println()
		} else {
			//find where is the substring in the following line
			sub:="i"
			reI := regexp.MustCompile(sub)
			indOfSub := reI.FindAllStringIndex(line, -1)
			// Create a slice to store the lines for the current row
			var lineSlice string
			// Iterate over the 8 rows for each line
			for i := 0; i < 8; i++ {
				// Iterate over the characters in the current line
				for j, ch := range line {
					// Check if the character is a valid ASCII character
					if ch < 32 || ch > 128 {
						fmt.Println("error input")
						return
					}
					// Calculate the line index for the current character
					lineIndex := 1 + int(ch-32)*9 + i
					// Append the corresponding string to the line slice
					if ContSub(indOfSub, j) {
						lineSlice += "\033[31m" + str[lineIndex] + "\033[0m"
					} else {
						lineSlice += str[lineIndex]
					}
				}
				// Join the lines in the slice and print the row
				fmt.Println(lineSlice)

				// Clear the line slice for the next row
				lineSlice = ""
			}
		}
	}
}
