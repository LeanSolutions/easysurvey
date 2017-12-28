package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var (
	r RouterInfo
)

// RouterInfo contains Details of Router
type RouterInfo struct {
	Router *httprouter.Router
}

func init() {
	r.Router = httprouter.New()
}

// Instance returns the router
func Instance() *httprouter.Router {
	return r.Router
}


// Get is a shortcut for router.Handle("GET", path, handle)
func Get(path string, fn http.HandlerFunc) {
	r.Router.GET(path, HandlerFunc(fn))
}