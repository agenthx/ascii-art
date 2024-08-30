package main

import (
	"fmt"
	"regexp"
)

func main() {
	input := "n"
	reN := regexp.MustCompile(`.`)
	fmt.Println(reN.MatchString(input))
}
