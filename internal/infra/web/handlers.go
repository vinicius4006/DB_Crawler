package web

import (
	model "db_crawler/internal/entity"
	"db_crawler/internal/infra/usecase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handlers struct {
	SiteUseCase    usecase.SiteUseCase
	MetaTagUseCase usecase.MetaTagUseCase
	WordUseCase    usecase.WordUseCase
}

func NewHandlers(siteUseCase *usecase.SiteUseCase, metaTagUseCase *usecase.MetaTagUseCase, wordUseCase *usecase.WordUseCase) *Handlers {
	return &Handlers{
		SiteUseCase:    *siteUseCase,
		MetaTagUseCase: *metaTagUseCase,
		WordUseCase:    *wordUseCase,
	}
}

func (h *Handlers) CreateSiteHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *Handlers) GetSitesHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *Handlers) CreateMetaTagsHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *Handlers) GetMetaTagsHandler(w http.ResponseWriter, r *http.Request) {

	var output any
	var err error
	siteid := r.URL.Query().Get("siteid")
	tag := r.URL.Query().Get("tag")
	fmt.Println(siteid)
	if len(siteid) > 0 {
		parse, _ := strconv.ParseUint(siteid, 10, 64)
		output, err = h.MetaTagUseCase.ExecuteFindBySiteID(parse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if len(tag) > 0 {
		output, err = h.MetaTagUseCase.ExecuteFindByTag(tag)
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

func (h *Handlers) CreateWordsHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *Handlers) GetWordsHandler(w http.ResponseWriter, r *http.Request) {

}
