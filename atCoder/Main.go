package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()
	test := strings.Split(input, " ")
	current, err := strconv.Atoi(test[0])
	if err != nil {
		fmt.Println("失敗")
	}
	fmt.Println(current - 3)
	i := 1
	
	for {
		value, err := strconv.Atoi(test[i])
		if err != nil {
			fmt.Println("失敗")
		}
		current = current - value
		if(current < 1){
			break;
		}
		if(current > 2){
			i = 1
			continue;
		}
		i++
	}

	if(i == 1){
		fmt.Println("F")
	} else if (i == 2){
		fmt.Println("M")
	}else{
		fmt.Println("T")
	}
	
}