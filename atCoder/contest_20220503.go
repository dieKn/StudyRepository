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
  var input, output string
  var value1,value2 int
  var result float64
  sc := bufio.NewScanner(os.Stdin)
  if sc.Scan() {
    input = sc.Text() 
  }
  arr1 := strings.Split(input, " ")
  if strings.Index(arr1[0], "F") != -1{
    value1 = fToNum(arr1[0])
  }
  if strings.Index(arr1[0], "B") != -1{
    value1 = bfToNum(arr1[0])
  }
  if strings.Index(arr1[1], "F") != -1{
    value2 = fToNum(arr1[1])
  }
  if strings.Index(arr1[1], "B") != -1{
    value2 = bfToNum(arr1[1])
  }
  if value1 > 0 && value2 < 0 || value1 < 0 && value2 > 0{
    result = math.Abs(float64(value1) - float64(value2)) -1
  }else{
    result = math.Abs(float64(value1) - float64(value2))
  }
  output = strconv.Itoa(int(result))
  fmt.Println(output) 
}

func bfToNum(value string) int {
  result,err := strconv.Atoi(strings.Split(value, "B")[1])
    if err != nil{
      fmt.Println("エラー")
    }
    return result * -1
}
func fToNum(value string) int {
  result,err := strconv.Atoi(strings.Split(value, "F")[0])
    if err != nil{
      fmt.Println("エラー")
    }
    return result
}