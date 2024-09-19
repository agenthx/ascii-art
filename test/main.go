package main

import (
	"fmt"
	"strings"
)

// import (
// 	"fmt"
// 	"os"
// 	piscine "piscine/functions"
// 	"strings"
// )

// 	banner := "standard.txt"
// 	file, err := os.ReadFile(banner)
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	str := strings.Split(string(file), "\n")
// 	piscine.AsciiFCOJ(str, "input", "in", "orange", "")
// 	piscine.AsciiFCOJ(str, "shush", "sh", "purple", "")
// 	piscine.AsciiFCOJ(str, "shush", "u", "blue", "")

func main() {
	h := "hi hello 7"
	k := strings.Split(h, " ")
	fmt.Println(k)
}
