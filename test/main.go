package main

import (
	"fmt"
	"regexp"
)

func main() {
	// var color string
	// reset := "\033[0m"

	// switch os.Args[1] {
	// case "gray":
	// 	color = "\033[30m"
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
	// func ContSub(ch, sub, line string, index int) bool {
	//take in the full line and look for where the substrings are
	//see the ch given and if its the right ch of the substring using the index
	/*lime is mine*/

	// return true
	// }
	sub := "ime"
	// ch := "l"
	line := "lime is mine kime"
	// index := 1

	reI := regexp.MustCompile(sub)
	intt := reI.FindAllStringIndex(line, -1)
	fmt.Println(intt[0][0])
	

}
