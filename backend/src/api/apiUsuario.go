package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"models"
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tratamentoErro(err, http.StatusBadRequest, w)
		return
	}

	var usuario *models.Usuario
	err = json.Unmarshal(body, usuario)
	if err != nil {
		tratamentoErro(err, http.StatusBadRequest, w)
		return
	}

	err = db.GravarUsuario(usuario)
	if err != nil {
		tratamentoErro(err, http.StatusBadRequest, w)
		return
	}

	body, err = json.Marshal(usuario)
	if err != nil {
		tratamentoErro(err, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func AlterarUsuarioAPI(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tratamentoErro(err, http.StatusBadRequest, w)
		return
	}

	var usuario *models.Usuario
	err = json.Unmarshal(body, usuario)
	if err != nil {
		tratamentoErro(err, http.StatusBadRequest, w)
		return
	}

	err = db.AlterarUsuario(usuario)
	if err != nil {
		tratamentoErro(err, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RemoverUsuarioAPI(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	codigo, err := strconv.Atoi(params["codigo"])
	if err != nil || codigo == 0 {
		tratamentoErro(err, http.StatusBadRequest, w)
		return
	}

	err = db.RemoverUsuario(codigo)
	if err != nil {
		tratamentoErro(err, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
