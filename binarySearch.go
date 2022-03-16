package main

import "fmt"


func main() {
  value := make([]int, 0, 60)
  cap := cap(value)
  for len := len(value); len < cap; len++ {
	  len++;
	value = append(value, len)
  }

  for _,val := range value {
	  fmt.Println(val)
  }
  test := binarySearch(33,value)
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
	  if(target > valueSlice[mid]){
		  left = mid
		  mid = (mid +(right - mid)) /2
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


