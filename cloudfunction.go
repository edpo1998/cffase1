package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/edpo1998/cffase1/api"
)

func main() {
	api.Connectdb()
	http.HandleFunc("/", index)
	http.HandleFunc("/vm", api.LogsHandleFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcom to Cloud Functions")
}
