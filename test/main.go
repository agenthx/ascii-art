package main

import (
	"fmt"
	"strings"
)

func main() {
	// var color string
	// func ContSub(ch, sub, line string, index int) bool {
	//take in the full line and look for where the substrings are
	//see the ch given and if its the right ch of the substring using the index
	/*lime is mine*/

	color := strings.TrimPrefix("--color=red", "--color=")
	fmt.Println(color)
}
