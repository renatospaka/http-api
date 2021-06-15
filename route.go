package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	Id    int    `json: "id"`
	Title string `json: "title"`
	Text  string `json: "text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{{Id: 1, Title: "Title 1", Text: "Text 1"}}
}

func upAndRunning(resp http.ResponseWriter, req *http.Request) {
	log.Println("Up and Running")
	fmt.Fprintln(resp, "Up and Running")
}

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{error: "Error marshalling the post array"}`))
		return 
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}
