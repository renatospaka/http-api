package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/renatospaka/golang-rest-api/entity"
	"github.com/renatospaka/golang-rest-api/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func upAndRunning(resp http.ResponseWriter, req *http.Request) {
	log.Println("Up and Running")
	fmt.Fprintln(resp, "Up and Running")
}

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{error: "Error getting the posts"}`))
		return 
	}
	
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func addPosts(resp http.ResponseWriter, req *http.Request) {
	var post entity.Post
	resp.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{error: "Error unmarshalling the request"}`))
		return 
	}
	
	post.ID = rand.Int63()
	repo.Save(&post)
	
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}