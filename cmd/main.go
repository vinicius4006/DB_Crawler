package main

import (
	model "db_crawler/internal/entity"
	infra "db_crawler/internal/infra/db"
	"db_crawler/internal/infra/repository"
	"db_crawler/internal/infra/usecase"
	"db_crawler/internal/infra/web"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	gorm, err := infra.NewDB()

	if err != nil {
		log.Panic("failed to connect database")
	}
	fmt.Println("OK")

	// Let's run the migrations to create tables:
	gorm.DB.Migrator().AutoMigrate(&model.Site{}, &model.MetaTag{}, &model.Word{})

	if !gorm.DB.Migrator().HasConstraint(&model.MetaTag{}, "Ref") && !gorm.DB.Migrator().HasConstraint(&model.Word{}, "Ref") {
		// Let's create the constrants
		gorm.DB.Migrator().CreateConstraint(&model.MetaTag{}, "Ref")
		gorm.DB.Migrator().CreateConstraint(&model.Word{}, "Ref")
	}

	// Integration
	siteRepository := repository.NewSiteRepositoryPostgres(gorm.DB)
	siteUseCase := usecase.NewSiteUseCase(siteRepository)
	//
	metaTagRepository := repository.NewMetaTagRepositoryPostgres(gorm.DB)
	metaTagUseCase := usecase.NewMetaTagUseCase(metaTagRepository)
	//
	wordRepository := repository.NewWordRepositoryPostgres(gorm.DB)
	wordUseCase := usecase.NewWordUseCase(wordRepository)
	//
	handlers := web.NewHandlers(siteUseCase, metaTagUseCase, wordUseCase)
	r := mux.NewRouter()

	//Sites Routers
	r.HandleFunc("/api/sites", handlers.Sites.Create).Methods("POST")
	r.HandleFunc("/api/sites", handlers.Sites.Update).Methods("PATCH")
	r.HandleFunc("/api/sites/{id}", handlers.Sites.Get).Methods("GET")
	r.HandleFunc("/api/sites", handlers.Sites.Get).Queries("url", "{url}").Methods("GET")

	// MetaTags Routers
	r.HandleFunc("/api/metatags", handlers.MetaTags.Create).Methods("POST")
	r.HandleFunc("/api/metatags", handlers.MetaTags.Get).Queries("siteid", "{siteid}", "tag", "{tag}").Methods("GET")

	// Words Routers
	r.HandleFunc("/api/words", handlers.Words.Create).Methods("POST")
	r.HandleFunc("/api/words", handlers.Words.Update).Methods("PATCH")
	r.HandleFunc("/api/words", handlers.Words.Get).Queries("siteid", "{siteid}", "value", "{value}").Methods("GET")
	err = http.ListenAndServe(":5050", r)

	if err != nil {
		fmt.Println("Server don't run")
	}

}
