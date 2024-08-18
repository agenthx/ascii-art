package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filee, err := os.ReadFile("standard.txt")

	file := string(filee)
	if err != nil {
		fmt.Println(err)
		return
	}

  // ascii:= '!'

  count := 0
	str := strings.Split(file, "\n")
	for i := 1; i < len(str); i++ {

		if i%9 == 0 {
			fmt.Println()
		} else {
      if i%8 == 0{
        count++
      }

			fmt.Println(str[i])
		}
	}
  //make a loop that prints the first line to the 8th line of letters together
  
  for i := 0; i < len(str); i++ {
    
  }


}
