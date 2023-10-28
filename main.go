package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

type Aluno struct {
	Pessoa
	RA int `json:"RA"`
}

func getAlunos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(alunos)
}

var alunos []Aluno

func printHelloWorld(w http.ResponseWriter, r *http.Request) {
	var retorno string = "Hello World!"
	json.NewEncoder(w).Encode(retorno)
}

func postAluno(w http.ResponseWriter, r *http.Request) {
	var aluno Aluno
	err := json.NewDecoder(r.Body).Decode(&aluno)
	if err != nil {
		return
	}
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", printHelloWorld).Methods("GET")
	router.HandleFunc("/aluno", printHelloWorld).Methods("GET")
	router.HandleFunc("/aluno", printHelloWorld).Methods("POST")
	http.Handle("/hello", router)
	fmt.Println("to rodante")
	http.ListenAndServe(":8000", nil)
}
