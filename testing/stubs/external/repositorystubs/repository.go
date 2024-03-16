package repositorystubs

import (
	"context"

	"github.com/ooaklee/ghatd-detail-api-demo-endpoints/external/dictionary"
)

type Repository struct {
	GetWordsResponse []dictionary.Word
	GetWordsError    error

	GetWordByIdResponse *dictionary.Word
	GetWordByIdError    error

	GetWordByNameResponse *dictionary.Word
	GetWordByNameError    error

	CreateWordResponse *dictionary.Word
	CreateWordError    error

	DeleteWordByIdError error

	CreateWordsResponse []dictionary.Word
	CreateWordsError    error
}

func (r *Repository) GetWordByName(ctx context.Context, name string) (*dictionary.Word, error) {
	return r.GetWordByNameResponse, r.GetWordByNameError
}

func (r Repository) DeleteWordById(ctx context.Context, id string) error {
	return r.DeleteWordByIdError
}

func (r Repository) CreateWord(ctx context.Context, name string) (*dictionary.Word, error) {
	return r.CreateWordResponse, r.CreateWordError
}

func (r Repository) GetWordById(ctx context.Context, id string) (*dictionary.Word, error) {
	return r.GetWordByIdResponse, r.GetWordByIdError
}

func (r Repository) GetWords(ctx context.Context) ([]dictionary.Word, error) {
	return r.GetWordsResponse, r.GetWordsError
}
