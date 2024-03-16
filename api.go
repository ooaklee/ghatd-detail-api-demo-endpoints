package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ooaklee/ghatd/external/response"
	"github.com/ooaklee/ghatd/external/router"
	"github.com/ooaklee/ghatd/external/validator"

	//>ghatd {{ define "ApiDetailImports" }}
	"github.com/ooaklee/ghatd-detail-api-demo-endpoints/external/dictionary"
	"github.com/ooaklee/ghatd-detail-api-demo-endpoints/external/repository"
	//>ghatd {{ end }}
)

// content holds our static web server content.
//
//>ghatd {{ define "ApiDetailEmbeds" }}

//go:embed external/web/ui/html/responses/*.tmpl.html
var content embed.FS

//>ghatd {{ end }}

const serverPort = ":4088"

func main() {

	tempRouterMiddlewares := []mux.MiddlewareFunc{}

	// Initialise validator
	appValidator := validator.NewValidator()

	// Initialise router
	httpRouter := router.NewRouter(response.GetResourceNotFoundError, response.GetDefault200Response, tempRouterMiddlewares...)

	// Prep web detail
	NewApiDetail(httpRouter, appValidator, content, "external/")

	// Define server
	srv := &http.Server{
		Addr:    serverPort,
		Handler: httpRouter.GetRouter(),
	}

	// Start listening
	fmt.Printf("\nServer is listening on port - %s\n", serverPort)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Default().Panicf("server/unable-to-start-server - %v", err)
	}
}

func NewApiDetail(httpRouter *router.Router, appValidator *validator.Validator, embeddedContent fs.FS, embeddedContentFilePathPrefix string) {

	//>ghatd {{ define "ApiDetailInit" }}

	// Initialise repository
	coreRepository := repository.NewInMememoryRepository()

	// TODO: Create Service(s)
	dictionaryService := dictionary.NewService(coreRepository)

	// TODO: Create Handler(s)
	dictionaryHandler := dictionary.NewHandler(dictionaryService, appValidator, embeddedContent, embeddedContentFilePathPrefix)

	// TODO: Attach package routes to router
	dictionary.AttachRoutes(&dictionary.AttachRoutesRequest{
		Router:  httpRouter,
		Handler: dictionaryHandler,
	})

	//>ghatd {{ end }}
}
