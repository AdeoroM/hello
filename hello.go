package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	post := Post{}
	postID := flag.String("postId", "1", "Save")
	flag.Parse()
	if err := getJson("https://jsonplaceholder.typicode.com/posts/"+*postID, &post); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("User ID:%v \nID: %v \nTitle: %v \nBody: %v \n\n\n", post.UserID, post.ID, post.Title, post.Body)

}
