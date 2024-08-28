package main

import (
	"fmt"
	"regexp"
)

func main() {
	sub := "i"
	line := "hi bie"
	reI := regexp.MustCompile(sub)
	indOfSub := reI.FindAllStringIndex(line, -1)
	fmt.Println(indOfSub[0][1],indOfSub[0][0], indOfSub)
}
