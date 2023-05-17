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

type HandlerMetaTags struct {
	MetaTagUseCase usecase.MetaTagUseCase
}

func (h *HandlerMetaTags) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var input model.MetaTag

	err = json.Unmarshal(body, &input)

	if err != nil {
		// Caso haja algum erro na conversão, retorna um erro
		http.Error(w, "Erro ao converter payload", http.StatusBadRequest)
		return
	}

	output, err := h.MetaTagUseCase.ExecuteCreate(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (h *HandlerMetaTags) Get(w http.ResponseWriter, r *http.Request) {

	var output any
	var err error
	siteid := r.URL.Query().Get("siteid")
	tag := r.URL.Query().Get("tag")
	fmt.Println(siteid)
	if len(siteid) > 0 && len(tag) == 0 {
		parse, _ := strconv.ParseUint(siteid, 10, 64)
		output, err = h.MetaTagUseCase.ExecuteFindBySiteID(parse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if len(tag) > 0 {
		var num uint64
		if len(siteid) > 0 {
			num, _ = strconv.ParseUint(siteid, 10, 64)
		}
		output, err = h.MetaTagUseCase.ExecuteFindByTag(tag, num)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		output, err = h.MetaTagUseCase.ExecuteFindAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
