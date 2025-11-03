package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := strings.TrimSpace(r.URL.Query().Get("name"))
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid Input"})
		return
	}

	firstChar := strings.ToUpper(string(name[0]))
	if firstChar >= "A" && firstChar <= "M" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{Message: "Hello " + name})
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid Input"})
}

func main() {
	http.HandleFunc("/hello-world", helloWorldHandler)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
