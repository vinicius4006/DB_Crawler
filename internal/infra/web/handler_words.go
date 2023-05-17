package web

import (
	model "db_crawler/internal/entity"
	"db_crawler/internal/infra/usecase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type HandlerWords struct {
	WordUseCase usecase.WordUseCase
}

func (h *HandlerWords) CreateWordsHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *HandlerWords) GetWordsHandler(w http.ResponseWriter, r *http.Request) {
	var output any
	var err error
	siteid := r.URL.Query().Get("siteid")
	value := r.URL.Query().Get("value")
	fmt.Println(siteid)
	if len(siteid) > 0 && len(value) == 0 {
		parse, _ := strconv.ParseUint(siteid, 10, 64)
		output, err = h.WordUseCase.ExecuteFindBySiteID(parse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if len(value) > 0 {
		var num uint64
		if len(siteid) > 0 {
			num, _ = strconv.ParseUint(siteid, 10, 64)
		}
		output, err = h.WordUseCase.ExecuteFindByValue(value, num)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		output, err = h.WordUseCase.ExecuteFindAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

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
