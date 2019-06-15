package main

import (
	"api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//r.Use() Adicionar middleware para verificar se usuário está logado
	r.HandleFunc("/api/usuario", api.ConsultarListaUsuarioAPI).Methods("GET")
	r.HandleFunc("/api/usuario/{codigo}", api.ConsultarUsuarioAPI).Methods("GET")
	r.HandleFunc("/api/usuario", api.CriarUsuarioAPI).Methods("POST")
	r.HandleFunc("/api/usuario", api.AlterarUsuarioAPI).Methods("PUT")
	r.HandleFunc("/api/usuario/{codigo}", api.RemoverUsuarioAPI).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
