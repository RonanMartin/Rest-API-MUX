package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/users", getUsers).Methods("GET") // aqui substituiu o http.HandleFunc("/users", getUsers) e acrescentou o método aceito.

	fmt.Println("api is on :8080")
	log.Fatal(http.ListenAndServe(":8080", muxRouter)) // muxRouter substituiu o "nil"
}

type User struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	//  if r.Method != "GET" {
	//	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	//	return
	//  }

	// "Toda parte acima foi substituida pelo	.Methods() que está na Func Main
	// e tem a função de informar quais os métodos que podem entrar ali"

	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			ID:   1,
			Name: "Ronan",
		},
		{
			ID:   2,
			Name: "José",
		},
	})
}
