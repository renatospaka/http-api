package controller

import (
	"encoding/json"
	"net/http"

	"github.com/renatospaka/golang-rest-api/entity"
	"github.com/renatospaka/golang-rest-api/service"
	errors "github.com/renatospaka/golang-rest-api/error"
)

type controller struct {}

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPosts(resp http.ResponseWriter, req *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the posts"})
		return 
	}
	
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (*controller) AddPosts(resp http.ResponseWriter, req *http.Request) {
	var post entity.Post
	resp.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling the post"})
		return 
	}
	
	errl := postService.Validate(&post)
	if errl != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: errl.Error()})
		return
	}
	
	newPost, errl := postService.Create(&post)
	if errl != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: errl.Error()})
		//json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(newPost)
}