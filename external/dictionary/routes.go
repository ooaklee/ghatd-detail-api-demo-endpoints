package dictionary

import (
	"fmt"
	"net/http"

	"github.com/ooaklee/ghatd/external/common"
	"github.com/ooaklee/ghatd/external/router"
)

// dictionaryHandler expected methods for valid dictionary handler
type dictionaryHandler interface {
	DeleteWord(w http.ResponseWriter, r *http.Request)
	CreateWord(w http.ResponseWriter, r *http.Request)
	GetWords(w http.ResponseWriter, r *http.Request)
	GetWordById(w http.ResponseWriter, r *http.Request)
}

const (
	// ApiDictionaryPrefix base Uri prefix for all dictionary routes
	ApiDictionaryPrefix = common.ApiV1UriPrefix

	// ApiWordsVariable Uri variable used to get actions words
	ApiWordsVariable = "/words"
)

var (
	// ApiDictionaryWordIdVariable Uri variable used to get dictionary Id out of Uri
	ApiDictionaryWordIdVariable = fmt.Sprintf("/{%s}", DictionaryWordUriVariableId)

	// ApiDictionarySpecificWordIdUriPath the Uri path for actioning a specific word
	ApiDictionarySpecificWordIdUriPath = fmt.Sprintf("%s%s", ApiWordsVariable, ApiDictionaryWordIdVariable)
)

// AttachRoutesRequest holds everything needed to attach dictionary
// routes to router
type AttachRoutesRequest struct {
	// Router main router being served by Api
	Router *router.Router

	// Handler valid dictionary handler
	Handler dictionaryHandler
}

// AttachRoutes attaches dictionary handler to corresponding
// routes on router
func AttachRoutes(request *AttachRoutesRequest) {
	httpRouter := request.Router.GetRouter()

	dictionaryRoutes := httpRouter.PathPrefix(ApiDictionaryPrefix).Subrouter()

	dictionaryRoutes.HandleFunc(ApiWordsVariable, request.Handler.GetWords).Methods(http.MethodGet, http.MethodOptions)
	dictionaryRoutes.HandleFunc(ApiWordsVariable, request.Handler.CreateWord).Methods(http.MethodPost, http.MethodOptions)
	dictionaryRoutes.HandleFunc(ApiDictionarySpecificWordIdUriPath, request.Handler.GetWordById).Methods(http.MethodGet, http.MethodOptions)
	dictionaryRoutes.HandleFunc(ApiDictionarySpecificWordIdUriPath, request.Handler.DeleteWord).Methods(http.MethodDelete, http.MethodOptions)

}
