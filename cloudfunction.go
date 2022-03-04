package main

import (
	"net/http"
	"os"

	"github.com/edpo1998/cffase1/api"
)

func main() {
	// Conect with mongo
	api.Connectdb()
	// Router
	http.HandleFunc("/", api.WelcomHandleFunc)
	http.HandleFunc("/vm", api.LogsHandleFunc)
	// Start Server
	http.ListenAndServe(port(), nil)
}

// Find enviroment variable port if return none then set default port
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	return ":" + port
}
