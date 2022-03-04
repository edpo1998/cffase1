package api

import (
	"encoding/json"
	"net/http"
)

// Welcom response structure
type Welcom struct {
	Message string
}

// WelcomHandleFunc to be used as http.HandleFunc for Hello API
func WelcomHandleFunc(w http.ResponseWriter, r *http.Request) {

	m := Welcom{"Welcome to Cloud Function"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
