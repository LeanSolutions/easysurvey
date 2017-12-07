package router

import "github.com/julienschmidt/httprouter"

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
