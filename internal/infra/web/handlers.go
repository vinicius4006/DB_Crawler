package web

import (
	"db_crawler/internal/infra/usecase"
)

type Handlers struct {
	Sites    HandlerSites
	MetaTags HandlerMetaTags
	Words    HandlerWords
}

func NewHandlers(siteUseCase *usecase.SiteUseCase, metaTagUseCase *usecase.MetaTagUseCase, wordUseCase *usecase.WordUseCase) *Handlers {
	return &Handlers{
		Sites:    HandlerSites{SiteUseCase: *siteUseCase},
		MetaTags: HandlerMetaTags{MetaTagUseCase: *metaTagUseCase},
		Words:    HandlerWords{WordUseCase: *wordUseCase},
	}
}
