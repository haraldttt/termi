package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Hello welcome to termi")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("You entered:",input)
}
