package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	fmt.Print(string(out))
	if err != nil {
		fmt.Println("Error: ", err)
	}
	hi:=strings.TrimSpace(string(out))
	fmt.Println(strconv.Atoi(hi))
}
