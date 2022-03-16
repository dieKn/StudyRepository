package main

import "fmt"


func main() {
  value := make([]int, 0, 60)
  cap := cap(value)
  for len := len(value); len < cap; len++ {
	  len+=2;
	value = append(value, len)
  }

  test := binarySearch(5,value)
  fmt.Println(test)  
}

func binarySearch(target int, valueSlice []int) int{
  length := len(valueSlice)
  mid := length /2
  right := length
  left := 0
  for left <= right {
	  if(valueSlice[mid] == target){
		  return mid
	  }
	  if(left +1 == right){
		  fmt.Println("not found")
		  break
	  }
	  if(target > valueSlice[mid]){
		  left = mid
		  mid = (right + mid) /2
		  continue
	  }
	  if(target < valueSlice[mid]){
		  right = mid
		  mid = (right - left) /2
		  continue
	  }
	  return -1
  }
  return 0;
}


