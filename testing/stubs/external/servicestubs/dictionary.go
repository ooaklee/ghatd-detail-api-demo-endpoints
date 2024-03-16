package servicestubs

import (
	"context"

	"github.com/ooaklee/ghatd-detail-api-demo-endpoints/external/dictionary"
)

type Dictionary struct {
	GetWordsResponse *dictionary.GetWordsResponse
	GetWordsError    error

	GetWordByIdResponse *dictionary.GetWordByIdResponse
	GetWordByIdError    error

	CreateWordResponse *dictionary.CreateWordResponse
	CreateWordError    error

	DeleteWordByIdError error
}

func (d *Dictionary) DeleteWordById(ctx context.Context, req *dictionary.DeleteWordRequest) error {
	return d.DeleteWordByIdError
}

func (d *Dictionary) CreateWord(ctx context.Context, req *dictionary.CreateWordRequest) (*dictionary.CreateWordResponse, error) {
	return d.CreateWordResponse, d.CreateWordError
}

func (d *Dictionary) GetWordById(ctx context.Context, req *dictionary.GetWordByIdRequest) (*dictionary.GetWordByIdResponse, error) {
	return d.GetWordByIdResponse, d.GetWordByIdError
}

func (d *Dictionary) GetWords(ctx context.Context, req *dictionary.GetWordsRequest) (*dictionary.GetWordsResponse, error) {
	return d.GetWordsResponse, d.GetWordsError
}
