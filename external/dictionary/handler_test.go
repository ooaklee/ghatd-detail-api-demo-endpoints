package dictionary_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ooaklee/ghatd-detail-api-demo-endpoints/external/dictionary"
	"github.com/ooaklee/ghatd-detail-api-demo-endpoints/testing/stubs/external/servicestubs"
	"github.com/ooaklee/ghatd/external/validator"
	responsehelpers "github.com/ooaklee/ghatd/testing/helpers"
	"github.com/ooaklee/reply"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetWordById(t *testing.T) {

	// Create struct to emulate shape of embedded error items
	type embeddedError struct {
		// Title a short summary of the problem
		Title string `json:"title,omitempty"`

		// Detail a description of the error
		Detail string `json:"detail,omitempty"`

		// About holds the link that gives further insight into the error
		About string `json:"about,omitempty"`

		// Status the HTTP status associated with error
		Status string `json:"status,omitempty"`

		// Code internal error code used to reference error
		Code string `json:"code,omitempty"`

		// Meta contains additional meta-information about the error
		Meta interface{} `json:"meta,omitempty"`
	}

	// Create struct to emulate shape of error response
	type errorResponse struct {
		Errors []embeddedError `json:"errors"`
	}

	tests := []struct {
		name               string
		dictionaryService  *servicestubs.Dictionary
		request            *http.Request
		assertResponse     func(w *httptest.ResponseRecorder, t *testing.T)
		expectedStatusCode int
		expectedMessage    string
	}{
		{
			name: "Success - Word found",
			dictionaryService: &servicestubs.Dictionary{
				GetWordByIdResponse: &dictionary.GetWordByIdResponse{
					Word: &getMockWords()[0],
				},
			},
			request: httptest.NewRequest(http.MethodGet, "/dictionary/words/8ba655eb-bcc6-4246-9c78-ac070cf3ac8e", nil),
			assertResponse: func(w *httptest.ResponseRecorder, t *testing.T) {

				embeddedResponse := dictionary.Word{}

				res := reply.NewResponseRequest{
					Data: &embeddedResponse,
				}

				err := responsehelpers.UnmarshalResponseBody(w, &res)
				if err != nil {
					t.Fatalf("GetWordById() failed, cannot get res content: %v", err)
				}

				expectedBody := dictionary.Word{Id: "8ba655eb-bcc6-4246-9c78-ac070cf3ac8e", Name: "fire truck", CreatedAt: "2021-04-01T15:04:05"}

				assert.Equal(t, &expectedBody, res.Data)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Failure - Word not found",
			dictionaryService: &servicestubs.Dictionary{
				GetWordByIdError: errors.New(dictionary.ErrKeyWordWithIdNotFound),
			},
			request: httptest.NewRequest(http.MethodGet, "/dictionary/words/bd2cbad1-6ccf-48e3-bb92-bc9961bc011e", nil),
			assertResponse: func(w *httptest.ResponseRecorder, t *testing.T) {

				res := errorResponse{}

				err := responsehelpers.UnmarshalResponseBody(w, &res)
				if err != nil {
					t.Fatalf("Cannot get response content: %v", err)
				}

				assert.Equal(t, errorResponse{Errors: []embeddedError{{Title: "Resource Not Found", Detail: "No word can be found matching the Id provided", About: "", Status: "404", Code: "R-002"}}}, res)

			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:              "Failure - Id validation failure",
			dictionaryService: &servicestubs.Dictionary{},
			request:           httptest.NewRequest(http.MethodGet, "/dictionary/words/incorrect-uuid-4", nil),
			assertResponse: func(w *httptest.ResponseRecorder, t *testing.T) {

				res := errorResponse{}

				err := responsehelpers.UnmarshalResponseBody(w, &res)
				if err != nil {
					t.Fatalf("Cannot get response content: %v", err)
				}

				assert.Equal(t, errorResponse{Errors: []embeddedError{{Title: "Bad Request", Detail: "Invalid or malformatted word identifier provided", About: "", Status: "400", Code: "R-005"}}}, res)

			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			v := validator.NewValidator()

			router := mux.NewRouter()

			// TODO: Investigate best way to test htmx integration
			router.HandleFunc("/dictionary/words/{wordId}", dictionary.NewHandler(test.dictionaryService, v, nil, "").GetWordById)
			router.ServeHTTP(w, test.request)

			test.assertResponse(w, t)
			assert.Equal(t, test.expectedStatusCode, w.Code)

		})
	}
}

func getMockWords() []dictionary.Word {

	return []dictionary.Word{
		{
			Id:        "8ba655eb-bcc6-4246-9c78-ac070cf3ac8e",
			Name:      "fire truck",
			CreatedAt: "2021-04-01T15:04:05",
		},
		{
			Id:        "8dcee940-7b91-4191-96c0-c14fb1f874af",
			Name:      "rubbish truck",
			CreatedAt: "2021-04-02T16:04:05",
		},
		{
			Id:        "58fec080-90b6-4cba-982b-2cdcc3997ae2",
			Name:      "car",
			CreatedAt: "2021-04-03T16:05:05",
		},
		{
			Id:        "553eb446c-4923-42fe-9d0c-19a209159631",
			Name:      "bed",
			CreatedAt: "2021-04-03T16:06:05",
		},
	}
}
