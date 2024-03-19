package dictionary

import (
	"context"
	"errors"
	"fmt"

	"github.com/ooaklee/ghatd/external/logger"
	"github.com/ooaklee/ghatd/external/toolbox"
)

// dictionaryRespository expected methods of a valid dictionary repository
type dictionaryRespository interface {
	GetWords(ctx context.Context) ([]Word, error)
	CreateWord(ctx context.Context, word string) (*Word, error)
	GetWordById(ctx context.Context, id string) (*Word, error)
	GetWordByName(ctx context.Context, name string) (*Word, error)
	DeleteWordById(ctx context.Context, id string) error
}

// Service holds and manages dictionary business logic
type Service struct {
	dictionaryRespository dictionaryRespository
}

// NewService created dictionary service
func NewService(dictionaryRespository dictionaryRespository) *Service {
	return &Service{
		dictionaryRespository: dictionaryRespository,
	}
}

// DeleteWordById attempts to delete the word with matching Id in repository
func (s *Service) DeleteWordById(ctx context.Context, r *DeleteWordRequest) error {

	logger := logger.AcquireFrom(ctx)

	logger.Info(fmt.Sprintf("checking-for-word-to-delete-with-id: %s", r.Id))
	_, err := s.dictionaryRespository.GetWordById(ctx, r.Id)
	if err != nil {

		return err
	}

	logger.Debug(fmt.Sprintf("found-word-with-id: %s", r.Id))

	logger.Info(fmt.Sprintf("requesting-repository-deletes-word-with-id: %s", r.Id))
	return s.dictionaryRespository.DeleteWordById(ctx, r.Id)
}

// CreateWord attempt to create the word in the repository
func (s *Service) CreateWord(ctx context.Context, r *CreateWordRequest) (*CreateWordResponse, error) {

	logger := logger.AcquireFrom(ctx)

	normalisedName := toolbox.StringStandardisedToLower(r.Name)

	logger.Info(fmt.Sprintf("creating-new-word-entry: %s", normalisedName))

	_, err := s.dictionaryRespository.GetWordByName(ctx, normalisedName)
	if err == nil {
		logger.Warn(fmt.Sprintf("word-entry-with-name-exists: %s", normalisedName))
		return nil, errors.New(ErrKeyWordAlreadyExists)
	}

	word, err := s.dictionaryRespository.CreateWord(ctx, normalisedName)
	if err != nil {
		return nil, err
	}

	logger.Info(fmt.Sprintf("sucessfully-created-new-entry-for-word: %s (%s)", word.Name, word.Id))

	return &CreateWordResponse{
		Word: word,
	}, nil
}

// GetWords returns the words from the repository
func (s *Service) GetWords(ctx context.Context, r *GetWordsRequest) (*GetWordsResponse, error) {
	logger := logger.AcquireFrom(ctx)

	logger.Info("requesting-all-words-in-repository")

	words, err := s.dictionaryRespository.GetWords(ctx)
	if err != nil {
		logger.Warn("failed-to-get-all-word-entries")
		return nil, err
	}

	return &GetWordsResponse{
		Words: words,
	}, nil
}

// GetWordById returns word with matching Id from repository
func (s *Service) GetWordById(ctx context.Context, r *GetWordByIdRequest) (*GetWordByIdResponse, error) {
	logger := logger.AcquireFrom(ctx)

	logger.Info(fmt.Sprintf("checking-for-word-to-get-with-id: %s", r.Id))
	word, err := s.dictionaryRespository.GetWordById(ctx, r.Id)
	if err != nil {

		return nil, err
	}

	logger.Info(fmt.Sprintf("sucessfully-found-entry-with-id: %s ", word.Id))
	return &GetWordByIdResponse{
		Word: word,
	}, nil
}
