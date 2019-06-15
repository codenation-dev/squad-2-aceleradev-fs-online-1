package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"db"

	"github.com/gorilla/mux"
)

func tratamentoErro(erro error, statusCode int, w http.ResponseWriter) {
	log.Fatal(erro)
	w.WriteHeader(statusCode)
}

func ConsultarUsuarioAPI(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	codigoBuscado, err := strconv.Atoi(params["codigo"])
	if err != nil || codigoBuscado == 0 {
		tratamentoErro(err, http.StatusBadRequest, w)
		return
	}

	usuario, err := db.ObtenhaUsuario(codigoBuscado)
	if err != nil {
		tratamentoErro(err, http.StatusInternalServerError, w)
		return
	}

	resposta, err := json.Marshal(usuario)
	if err != nil {
		tratamentoErro(err, http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resposta)
}

func ConsultarListaUsuarioAPI(w http.ResponseWriter, r *http.Request) {

	usuarios, err := db.ObtenhaListaUsuario()
	if err != nil {
		tratamentoErro(err, http.StatusInternalServerError, w)
		return
	}

	resposta, err := json.Marshal(usuarios)
	if err != nil {
		tratamentoErro(err, http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resposta)
}

func CriarUsuarioAPI(w http.ResponseWriter, r *http.Request) {
	teste := []byte("CriarUsuarioAPI")
	w.Write(teste)
}

func AlterarUsuarioAPI(w http.ResponseWriter, r *http.Request) {
	teste := []byte("AlterarUsuarioAPI")
	w.Write(teste)
}

func RemoverUsuarioAPI(w http.ResponseWriter, r *http.Request) {
	teste := []byte("RemoverUsuarioAPI")
	w.Write(teste)
}
