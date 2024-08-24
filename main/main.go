package main

import (
	"fmt"
	"os"
	"piscine"
	"strings"
)

func main() {
	//check if the arguments given are not exactly 2
	if len(os.Args) < 2 {
		fmt.Println("error number of arguments")
		return
	}
	//store the string given in the argument
	input := os.Args[1]
	//the default is standard style and if otherwise do a switch case
	banner := "standard.txt"
	if len(os.Args) == 3 {
		switch os.Args[2] {
		case "shadow":
			banner = "shadow.txt"
		case "thinkertoy":
			banner = "thinkertoy.txt"
		case "standard":
			break
		default:
			fmt.Println("Error: templates are 'standard' 'shadow' 'thinkertoy'")
			return
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
	piscine.PrintASCIIfs(str, input)
}
