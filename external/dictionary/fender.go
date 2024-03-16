package dictionary

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ooaklee/ghatd/external/logger"
	"github.com/ooaklee/ghatd/external/toolbox"
)

// mapRequestToGetWordByIdRequest maps incoming request for getting word by id to correct
// struct.
func mapRequestToGetWordByIdRequest(request *http.Request, validator dictionaryValidator) (*GetWordByIdRequest, error) {
	logger := logger.AcquireFrom(request.Context())

	parsedRequest := &GetWordByIdRequest{}

	wordId, err := getWordIdFromUri(request)
	if err != nil {
		return nil, err
	}

	parsedRequest.Id = wordId

	if err := validator.Validate(parsedRequest); err != nil {
		logger.Warn("validation-error-while-fetching-word-id-for-get-by-id-request")
		return nil, errors.New(ErrKeyInvalidWordId)
	}

	return parsedRequest, nil
}

// mapRequestToDeleteWordRequest maps incoming delete word request to correct
// struct.
func mapRequestToDeleteWordRequest(request *http.Request, validator dictionaryValidator) (*DeleteWordRequest, error) {

	parsedRequest := &DeleteWordRequest{}

	wordId, err := getWordIdFromUri(request)
	if err != nil {
		return nil, err
	}

	parsedRequest.Id = wordId

	if err := validator.Validate(parsedRequest); err != nil {
		return nil, errors.New(ErrKeyInvalidWordId)
	}

	return parsedRequest, nil
}

// mapRequestToCreateWordRequest maps incoming create word request to correct
// struct.
func mapRequestToCreateWordRequest(request *http.Request, validator dictionaryValidator) (*CreateWordRequest, error) {
	parsedRequest := &CreateWordRequest{}

	err := toolbox.DecodeRequestBody(request, parsedRequest)
	if err != nil {
		return nil, errors.New(ErrKeyInvalidWordBody)
	}

	if err := validator.Validate(parsedRequest); err != nil {
		return nil, errors.New(ErrKeyInvalidWordBody)
	}

	return parsedRequest, nil
}

// mapRequestToGetWordsRequest maps incoming GetWords request to correct
// struct.
func mapRequestToGetWordsRequest(request *http.Request, validator dictionaryValidator) (*GetWordsRequest, error) {
	parsedRequest := &GetWordsRequest{}

	if err := validator.Validate(parsedRequest); err != nil {
		return nil, errors.New(ErrKeyInvalidQueryParam)
	}

	return parsedRequest, nil
}

// getWordIdFromUri pulls word Id from Uri. If fails, returns error
func getWordIdFromUri(request *http.Request) (string, error) {
	var wordId string

	if wordId = mux.Vars(request)[DictionaryWordUriVariableId]; wordId == "" {
		return "", errors.New(ErrKeyInvalidWordId)
	}

	return wordId, nil
}
