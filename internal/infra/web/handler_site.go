package web

import (
	model "db_crawler/internal/entity"
	"db_crawler/internal/infra/usecase"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HandlerSites struct {
	SiteUseCase usecase.SiteUseCase
}

func (h *HandlerSites) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var input model.Site

	err = json.Unmarshal(body, &input)

	if err != nil {
		// Caso haja algum erro na conversão, retorna um erro
		http.Error(w, "Erro ao converter payload", http.StatusBadRequest)
		return
	}
	output, err := h.SiteUseCase.ExecuteCreate(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (h *HandlerSites) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var output any
	var err error
	url := r.URL.Query().Get("url")

	if v, ok := vars["id"]; ok {
		parse, _ := strconv.ParseUint(v, 10, 64)
		output, err = h.SiteUseCase.ExecuteFindByID(parse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if len(url) > 0 {

		output, err = h.SiteUseCase.ExecuteFindByURL(url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		output, err = h.SiteUseCase.ExecuteFindAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *HandlerSites) Update(w http.ResponseWriter, r *http.Request) {

	var updateSite model.Site
	err := json.NewDecoder(r.Body).Decode(&updateSite)

	if err != nil {
		// Caso haja algum erro na conversão, retorna um erro
		http.Error(w, "Erro ao converter payload", http.StatusBadRequest)
		return
	}

	err = h.SiteUseCase.ExecuteUpdate(&updateSite)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
