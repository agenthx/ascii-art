package piscine

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func width() int {
	cmd := exec.Command("tput", "cols") //the command to check width of terminal
	cmd.Stdin = os.Stdin                //sets the standard input of the os to be same for output
	out, err := cmd.Output()            //run the command and store the output in out
	if err != nil {
		fmt.Println("Error: ", err)
	}
	lenOfT := strings.TrimSpace(string(out)) //remove extra spaces from the string out so we just get the digits
	lt, _ := strconv.Atoi(lenOfT)            //convert string to int
	return lt
}
