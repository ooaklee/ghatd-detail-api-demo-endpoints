package dictionary

import (
	"reflect"

	"github.com/ooaklee/ghatd/external/toolbox"
)

type Word struct {
	// Id the unique identifier of the word
	Id string `json:"id"`

	// Name the word itself
	Name string `json:"name"`

	// CreateAt the time the word was created
	CreatedAt string `json:"created_at,omitempty"`
}

// SetCreatedAtTimeToNow sets the createdAt time to now (UTC)
func (w *Word) SetCreatedAtTimeToNow() *Word {
	w.CreatedAt = toolbox.TimeNowUTC()
	return w
}

func (w *Word) GetEntityTypeName() string {
	return toolbox.StringStandardisedToLower("word")
}
func (w *Word) GetEntityType() reflect.Type {
	return reflect.TypeOf(w)
}
