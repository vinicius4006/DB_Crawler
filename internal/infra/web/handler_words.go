package web

import (
	model "db_crawler/internal/entity"
	"db_crawler/internal/infra/usecase"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HandlerWords struct {
	WordUseCase usecase.WordUseCase
}

func (h *HandlerWords) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var input model.Word

	err = json.Unmarshal(body, &input)

	if err != nil {
		// Caso haja algum erro na conversão, retorna um erro
		http.Error(w, "Erro ao converter payload", http.StatusBadRequest)
		return
	}

	output, err := h.WordUseCase.ExecuteCreate(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (h *HandlerWords) Get(w http.ResponseWriter, r *http.Request) {
	var output any
	var err error
	q := r.URL.Query().Get("q")

	output, err = h.WordUseCase.ExecuteFindByValues(q)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *HandlerWords) Update(w http.ResponseWriter, r *http.Request) {

	var updateWord model.Word
	err := json.NewDecoder(r.Body).Decode(&updateWord)

	if err != nil {
		// Caso haja algum erro na conversão, retorna um erro
		http.Error(w, "Erro ao converter payload", http.StatusBadRequest)
		return
	}

	err = h.WordUseCase.ExecuteUpdate(&updateWord)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
