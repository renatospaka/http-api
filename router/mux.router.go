package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	//fmt.Printf("Mux HTTP Server listening on port ", port)
	log.Printf("Mux HTTP Server listening on port %s", port)

	http.ListenAndServe(port, muxDispatcher)
}

