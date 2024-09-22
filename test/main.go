package main

import (
	"fmt"
	"os"
	piscine "piscine/functions"
	"strings"
)

func main() {
	banner := "standard.txt"
	file, err := os.ReadFile("banner/" + banner)
	if err != nil {
		fmt.Print(err)
	}
	str := strings.Split(string(file), "\n")
	piscine.AsciiFCOJ(str, "hahah\nhahhah", "", "", "", "center")
	piscine.AsciiFCOJ(str, "1 two 4", "", "", "", "justify")
	piscine.AsciiFCOJ(str, "23/32", "", "", "", "left")
}
