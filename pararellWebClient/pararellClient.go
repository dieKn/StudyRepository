package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Item struct {
        Title     string    `json:"title"`
        CreatedAt time.Time `json:"created_at"`
}

func main() {
	results := []Item{}
	ch1 := make(chan []Item)
	ch2 := make(chan []Item)
	go execApi("1つ目","https://qiita.com/api/v2/users/snaka/items?page=3&per_page=10", ch1)
	go execApi("2つ目","https://qiita.com/api/v2/users/snaka/items?page=2&per_page=10", ch2)
	results = append(results, <-ch1...)
	results = append(results, <-ch2...)
	for i,result := range results {
		fmt.Printf("%d: %s %s\n",i,result.CreatedAt,result.Title)
	}
}

func execApi(title string, url string, ch chan []Item){
	resp, err := http.Get(url)
	if err != nil {
			log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
			log.Fatal(err)
	}

	var data []Item 

	if err := json.Unmarshal(body, &data); err != nil {
			log.Fatal(err)
	}
	fmt.Println(title+ "終了") 
	ch <- data
}