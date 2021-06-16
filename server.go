package main

import (
	"log"
	"net/http"

	"github.com/renatospaka/golang-rest-api/router"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	const port = ":8000"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(w, "Up and Running")	
	})

	// httpRouter.GET("/posts", GetPosts) 
	// httpRouter.POST("/posts", AddPosts) 

	httpRouter.SERVE(port)
}