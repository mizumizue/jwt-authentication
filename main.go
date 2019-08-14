package main

import (
	"encoding/json"
	"github.com/trewanek/jwt-authentication/middleware/auth"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()
	r.Handle("/auth", http.HandlerFunc(auth.GetJwtTokenHandler))
	r.Handle("/", auth.JwtMiddleware(http.HandlerFunc(rootHandler)))

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", nil)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(&Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Hello World",
	})
}
