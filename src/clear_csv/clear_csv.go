package main // has to be clear_csv

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
  fmt.Print("get csv > ")
  text, _ := reader.ReadString('\n')
	fmt.Printf(text)
}