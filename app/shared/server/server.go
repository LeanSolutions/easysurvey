package server

import (
	"fmt"
	"net/http"
	"time"
)

// Server stores all the information about server
type Server struct {
	Hostname string `json:"Hostname"` // Server Name
	UseHTTP  bool   `json:"UseHTTP"`  // Listen on HTTP
	HTTPPort int    `json:"HTTPPort"` // HTTP port
}

// Run will start http handler
func Run(httpHandlers http.Handler, s Server) {
	if s.UseHTTP {
		startHTTP(httpHandlers, s)
	}
}

// startHTTP starts the HTTP listener
func startHTTP(handlers http.Handler, s Server) {
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTP "+httpAddress(s))

	// Start the HTTP listener
	http.ListenAndServe(httpAddress(s), handlers)
}

// httpAddress returns the HTTP address
func httpAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.HTTPPort)
}
