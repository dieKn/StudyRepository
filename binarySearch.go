package main

import "fmt"


func main() {
  value := make([]int, 0, 101)
  target := 0
  cap := cap(value)
  for len := len(value); len < cap; len++ {
	  value = append(value, len)
	  len++;
	  len++;
  }

  answerIndex := binarySearchByRecursion(value,target,60,0, len(value)/2)
  fmt.Println(answerIndex)
  if(answerIndex > 0){
	  fmt.Println(value[answerIndex] == target)
  }
}

func binarySearchByRecursion(valueSlice []int, target int, right int, left int, mid int) int{
  if(right > left && mid != left && mid != right){
	  if(valueSlice[mid] == target){
		  return mid
	  }
	  if(valueSlice[mid] < target){
		  return binarySearchByRecursion(valueSlice, target, right, mid,(right + mid)/2)
		}
		return binarySearchByRecursion(valueSlice, target, mid, left, (mid - left)/2)
	}
  return -1;
}
func binarySearchByRoop(target int, valueSlice []int) int{
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


