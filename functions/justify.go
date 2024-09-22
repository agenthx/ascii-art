package piscine

import (
	"fmt"
	"os"
)

func Justify(str, lines []string, noOfS, noSpaceLength []int) {
	// Create a slice to store the lines for the current row
	var lineSlice string
	var space string
	var flag bool
	// Iterate over the lines
	for i, line := range lines {
		if line == "" {
			fmt.Println()
		} else {
			space, flag = Ss(i, noOfS[i], noSpaceLength[i])
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
				if flag {
					fmt.Println("|" + lineSlice + " |")
				} else {
					fmt.Println("|" + lineSlice + "|")
				}
				lineSlice = ""
			}
		}
	}
}

func Ss(i, noOfS, noSpaceLength int) (string, bool) {
	var space string
	flag := false
	if noOfS == 0 {
		return "", flag
	}
	fmt.Println(width(), noSpaceLength)
	hi := width() - noSpaceLength
	fmt.Println(hi, (hi/noOfS))
	if hi%2 == 1 {
		flag = true
	}
	for j := 1; j < hi/noOfS; j++ {
		space += " "
	}
	fmt.Println(flag,len(space))
	return space, flag
}
