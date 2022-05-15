package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		input = sc.Text()
	}
	arr := strings.Split(input, " ")
	var arrNum []int
	var output float64
	for _, a := range arr {
		i, _ := strconv.Atoi(a)
		arrNum = append(arrNum, i)
	}
	if arrNum[0] > arrNum[1] {
		output = 0
	} else {
		output = math.Ceil(float64(arrNum[1]-arrNum[0]) / 10)
	}
	fmt.Println(output)

}
