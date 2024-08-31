package main

import (
	"fmt"
	"os"
	"piscine"
	"regexp"
	"strings"
)

func main() {
	//check if the arguments given are arguments are not too little or much
	if len(os.Args) < 2 || len(os.Args) > 5 {
		fmt.Println("error: no.of arguments")
		return
	}
	banner := "standard.txt"
	var color string
	var sub string
	var input string
	var fileN string
	b4Last := os.Args[len(os.Args)-2]
	lastArg := os.Args[len(os.Args)-1]
	//find the banner
	reF := regexp.MustCompile(`(standard|shadow|thinkertoy)`)
	//find if file is txt
	reT := regexp.MustCompile(`\.txt$`)
	//if the argument starts with color flag
	if strings.Contains(os.Args[1], "--color=") {
		//store the color name
		color = strings.TrimPrefix(os.Args[1], "--color=")
		//if the last arg is a banner then b4 it is the string
		if reF.MatchString(lastArg) {
			banner = lastArg + ".txt"
			input = b4Last
		} else {
			if len(os.Args) >= 5 {
				fmt.Println("Error: arguments given are incorrect")
				os.Exit(1)
			}
			input = lastArg
		}
		//if arg 2 is differnt than input put it as sub
		if input == os.Args[2] {
			sub = input
		} else {
			sub = os.Args[2]
		}
	} else if strings.Contains(os.Args[1], "--output=") {
		fileN = strings.TrimPrefix(os.Args[1], "--output=")
		if reF.MatchString(fileN) || !reT.MatchString(fileN) {
			fmt.Println("Error: File name invalid!")
			os.Exit(3)
		}
		input = os.Args[2]
		if len(os.Args) == 4 {
			banner = os.Args[3] + ".txt"
			if !reF.MatchString(os.Args[3]) {
				fmt.Println("banner name incorrect")
				os.Exit(4)
			}
		}
	} else {
		if len(os.Args) >= 4 {
			fmt.Println("Error: arguments given are incorrect")
			os.Exit(2)
		}
		input = os.Args[1]
		if reF.MatchString(lastArg) && len(os.Args) > 2 {
			banner = os.Args[2] + ".txt"
		}
	}
	//open and store the contents of the standard.txt file
	file, err := os.ReadFile(banner)
	if err != nil {
		fmt.Print(err)
	}
	//split the contents of standard.txt file into line by line and store each line in a slice in the array
	str := strings.Split(string(file), "\n")
	//give the input and str into PrintASCII function
	piscine.AsciiFsColor(str, input, sub, color, fileN)
}
