package dictionary_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ooaklee/ghatd-detail-api-demo-endpoints/external/dictionary"
	"github.com/ooaklee/ghatd-detail-api-demo-endpoints/testing/stubs/external/repositorystubs"
	"github.com/stretchr/testify/assert"
)

func TestService_GetWordById(t *testing.T) {
	tests := []struct {
		name             string
		repository       *repositorystubs.Repository
		request          *dictionary.GetWordByIdRequest
		expectedResponse *dictionary.GetWordByIdResponse
		expectedError    error
	}{
		{
			name: "Success - Word found",
			repository: &repositorystubs.Repository{
				GetWordByIdResponse: &getMockWords()[0],
			},
			request: &dictionary.GetWordByIdRequest{
				Id: "8ba655eb-bcc6-4246-9c78-ac070cf3ac8e",
			},
			expectedResponse: &dictionary.GetWordByIdResponse{
				Word: &dictionary.Word{
					Id:        "8ba655eb-bcc6-4246-9c78-ac070cf3ac8e",
					Name:      "fire truck",
					CreatedAt: "2021-04-01T15:04:05",
				},
			},
			expectedError: nil,
		},
		{
			name: "Repository Error",
			repository: &repositorystubs.Repository{
				GetWordByIdError: errors.New("boom boom pow"),
			},
			request: &dictionary.GetWordByIdRequest{
				Id: "8ba655eb-bcc6-4246-9c78-ac070cf3ac8e",
			},
			expectedError: errors.New("boom boom pow"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := dictionary.NewService(test.repository)
			res, err := service.GetWordById(context.Background(), test.request)

			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedResponse, res)

		})
	}
}
