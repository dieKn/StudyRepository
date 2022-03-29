package main

import "fmt"

func main() {
  fmt.Printf("Hello World\n")
  v := 0
  // for _,v :=range(make([]int, 10)){
  for v <= 10{
    fmt.Println(v,`合計`,recursion(v))
    v++
  }
}


func recursion(val int) int{
  if(val == 0) {return 0;}
  return val + recursion(val -1);
}
