package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)
 
func main() {
  var input string
  sc := bufio.NewScanner(os.Stdin)
  if sc.Scan() {
    input = sc.Text()
  }
  i, _ := strconv.Atoi(input)
  if math.Pow(2,float64(i)) > math.Pow(float64(i),2) {
	  fmt.Println("Yes")
  } else {
	  fmt.Println("No")
  }
}