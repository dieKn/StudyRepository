package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		input = sc.Text()
	}
	fmt.Println("0" + input[:3])
}
