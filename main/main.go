package main

import (
	"fmt"
	"os"
	"piscine"
	"strings"
)

func main() {
	var input string
	var banner string
	//check if the arguments given are not exactly 2
	if len(os.Args) < 1 && len(os.Args) > 4 {
		fmt.Println("error number of arguments")
		return
	}
	if strings.Contains(os.Args[1], "--color=") {
		//input string
		input = os.Args[3]
		//color
		color := strings.TrimPrefix(os.Args[1], "--color=")
		// reset := "\033[0m"

		// switch os.Args[1] {
		// case "gray":
		// 	color = "\033[30mhi hello"
		// case "red":
		// 	color = "\033[31m"
		// case "green":
		// 	color = "\033[32m"
		// case "yellow":
		// 	color = "\033[33m"
		// case "blue":
		// 	color = "\033[34m"
		// case "purple":
		// 	color = "\033[35m"
		// case "cyan":
		// 	color = "\033[36m"
		// }

	} else {
		//store the string given in the argument
		input = os.Args[1]
		//the default is standard style and if otherwise do a switch case
		banner = "standard.txt"
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
