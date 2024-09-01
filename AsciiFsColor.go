package piscine

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func AsciiFsColor(str []string, input, sub, color string) {
	//if input is empty return nothing
	if input == "" {
		return
	}
	//check if color string is empty if not do case to match it
	var flag bool
	var reset string
	if color != "" {
		flag = true
		reset = "\033[0m"

		switch color {
		case "gray":
			color = "\033[30m"
		case "red":
			color = "\033[31m"
		case "green":
			color = "\033[32m"
		case "yellow":
			color = "\033[33m"
		case "blue":
			color = "\033[34m"
		case "purple":
			color = "\033[35m"
		case "cyan":
			color = "\033[36m"
		case "white":
			color = "\033[00m"
		default:
			fmt.Println("invalid color")
			os.Exit(3)
		}

	}

	// Replace all "\n" in the input string with an actual newline
	input2 := strings.ReplaceAll(input, "\\n", "\n")
	//check if the input consists of only \n
	reN := regexp.MustCompile(`[^\\n]`)
	if !reN.MatchString(input) && input != "n" {
		fmt.Print(input2)
		return
	}
	// Split the input string into lines
	lines := strings.Split(input2, "\n")
	//for sub things
	var sub2 []string
	if flag && sub == ""{
		sub = strings.ReplaceAll(sub, "\\n", "\n")
		sub2 = strings.Split(sub, "\n")
	}
	// Iterate over the lines
	for s, line := range lines {
		if line == "" {
			fmt.Println()
		} else {
			//find where is the substring in the following line
			var indOfSub [][]int
			if flag {
				reI := regexp.MustCompile(sub2[s])
				indOfSub = reI.FindAllStringIndex(line, -1)
			}
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
						lineSlice += color + str[lineIndex] + reset
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
