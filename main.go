package main

import (
	"fmt"
	"os"
	piscine "piscine/functions"
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
	var align string
	b4Last := os.Args[len(os.Args)-2]
	lastArg := os.Args[len(os.Args)-1]
	//find the banner
	reF := regexp.MustCompile(`(^standard$|^shadow$|^thinkertoy$)`)
	//find if file is txt
	reFF := regexp.MustCompile(`(^standard\.txt$|^shadow\.txt$|^thinkertoy\.txt$)`)
	//find if align is valid
	reA := regexp.MustCompile(`center|left|right|justify`)
	//if the argument starts with color flag
	if strings.Contains(os.Args[1], "--color=") {
		//store the color name
		color = strings.TrimPrefix(os.Args[1], "--color=")
		//if the last arg is a banner then b4 it is the string
		if reFF.MatchString(lastArg) || reF.MatchString(lastArg) {
			if reFF.MatchString(lastArg) {
				banner = lastArg
			} else {
				banner = lastArg + ".txt"
			}
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
		if len(os.Args) == 2 {
			fmt.Println("eror: no-of args")
			return
		}
		fileN = strings.TrimPrefix(os.Args[1], "--output=")
		if reF.MatchString(fileN) || reFF.MatchString(fileN) {
			fmt.Println("Error: File name invalid!")
			os.Exit(3)
		}
		input = os.Args[2]
	} else if strings.Contains(os.Args[1], "--align=") {
		if len(os.Args) <= 2 {
			fmt.Println("no enough args")
			os.Exit(6)
		}
		align = strings.TrimPrefix(os.Args[1], "--align=") //where to align text
		//check if the align is valid
		if !reA.MatchString(align) {
			fmt.Println("Error: aligns are: center, left, right, justify")
			os.Exit(5)
		}
		input = os.Args[2]
	} else {
		if len(os.Args) >= 4 {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --[OPTION]=<fileName.txt> something standard")
			os.Exit(2)
		}
		input = os.Args[1]
		if len(os.Args) > 2 {
			if reFF.MatchString(os.Args[2]) {
				banner = os.Args[2]
			} else if reF.MatchString(lastArg) {
				banner = os.Args[2] + ".txt"
			} else {
				fmt.Println("banner name incorrect")
				os.Exit(4)
			}
		}
	}
	//for banners
	if len(os.Args) == 4 {
		if reFF.MatchString(os.Args[3]) {
			banner = os.Args[3]
		} else if reF.MatchString(os.Args[3]) {
			banner = os.Args[3] + ".txt"
		} else {
			fmt.Println("banner name incorrect")
			os.Exit(4)
		}
	}
	//open and store the contents of the standard.txt file
	file, err := os.ReadFile("banner/" + banner)
	if err != nil {
		fmt.Print(err)
	}
	//split the contents of standard.txt file into line by line and store each line in a slice in the array
	str := strings.Split(string(file), "\n")
	//give the input and str into PrintASCII function
	piscine.AsciiFCOJ(str, input, sub, color, fileN, align)
}
