package atcoder_test

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)
 
func main() {
  var input, output string
  var results []string
  sc := bufio.NewScanner(os.Stdin)
  if sc.Scan() {
    input = sc.Text() 
  }
  if sc.Scan() {
    input = sc.Text() 
  }
  arr1 := strings.Split(input, " ")
  
  m := make(map[string]struct{})
  for _, ele := range arr1 {
    m[ele] = struct{}{}
  }
  uniq := [] string{}
  for i := range m {
    uniq = append(uniq, i)
  }
  for _,rv := range uniq {
	removed1 := strings.TrimSpace(strings.Replace(input, rv, "", -1))
	removed2 := strings.Replace(removed1, " ", "", -1)
	splited := strings.Split(removed2, "")
	joined := strings.Join(splited[:], " ")
	results = append(results, joined)
  }
  sort.Slice(results, func(i, j int) bool {
	return results[i] < results[j]
})
output = results[0]
fmt.Println(output)
}