package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/renatospaka/golang-rest-api/controller"
	"github.com/renatospaka/golang-rest-api/router"
)

var (
	httpRouter     router.Router             = router.NewChiRouter()
	postController controller.PostController = controller.NewPostController()
)

func main() {
	const port = ":8000"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Up and Running on port %s", port)
		fmt.Fprintln(w, "Up and Running on port ", port)
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPosts)

	httpRouter.SERVE(port)
}
