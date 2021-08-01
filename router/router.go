package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	GET(path string, f func(w http.ResponseWriter, r *http.Request))
	POST(path string, f func(w http.ResponseWriter, r *http.Request))
	PUT(path string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(path string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}

type MuxRouter struct{}

var (
	muxer = mux.NewRouter()
)

func CreateRouter() Router {
	return &MuxRouter{}
}

func (*MuxRouter) GET(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxer.HandleFunc(path, handler).Methods("GET")
}
func (*MuxRouter) POST(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxer.HandleFunc(path, handler).Methods("POST")
}
func (*MuxRouter) PUT(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxer.HandleFunc(path, handler).Methods("PUT")
}
func (*MuxRouter) DELETE(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxer.HandleFunc(path, handler).Methods("DELETE")
}
func (*MuxRouter) SERVE(port string) {
	http.ListenAndServe(port, muxer)
}
