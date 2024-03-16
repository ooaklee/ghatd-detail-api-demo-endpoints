package repository

import (
	"context"
	"fmt"

	"github.com/ooaklee/ghatd-detail-api-demo-endpoints/external/dictionary"
	"github.com/ooaklee/ghatd/external/toolbox"
)

// GetWords returns all the words in the repository
func (r *InMememoryRepository) GetWords(ctx context.Context) ([]dictionary.Word, error) {
	return r.Store, nil
}

// CreateWord adds word to the repository
func (r *InMememoryRepository) CreateWord(ctx context.Context, word string) (*dictionary.Word, error) {

	_, err := r.GetWordByName(ctx, word)
	if err == nil {
		return nil, fmt.Errorf(dictionary.ErrKeyWordAlreadyExists)
	}

	newWord := dictionary.Word{
		Name: word,
		Id:   toolbox.GenerateUuidV4(),
	}

	newWord.SetCreatedAtTimeToNow()

	r.Store = append(r.Store, newWord)

	return &newWord, nil
}

// GetWordById returns a word in the repository that matches the passed id
func (r *InMememoryRepository) GetWordById(ctx context.Context, id string) (*dictionary.Word, error) {

	for _, word := range r.Store {
		if word.Id == id {
			return &word, nil
		}
	}

	return nil, fmt.Errorf(dictionary.ErrKeyWordWithIdNotFound)
}

// GetWordById returns a word in the repository that matches the has the passed name
func (r *InMememoryRepository) GetWordByName(ctx context.Context, name string) (*dictionary.Word, error) {
	for _, word := range r.Store {
		if word.Name == name {
			return &word, nil
		}
	}

	return nil, fmt.Errorf(dictionary.ErrKeyWordWithNameNotFound)
}

// DeleteWordById removes a word in the repository that matches the has the passed id
func (r *InMememoryRepository) DeleteWordById(ctx context.Context, id string) error {

	for i, word := range r.Store {
		if word.Id == id {
			r.Store = removeWordFromStore(r.Store, i)
			return nil
		}
	}

	return fmt.Errorf(dictionary.ErrKeyWordWithIdNotFound)

}

// removeWordFromStore handles logic of removing the word at the passed index
// and returning an new slice
func removeWordFromStore(words []dictionary.Word, i int) []dictionary.Word {
	words[i] = words[len(words)-1]
	return words[:len(words)-1]
}
