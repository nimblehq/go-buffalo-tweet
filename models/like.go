package models

import (
    "encoding/json"
    "time"

    "github.com/markbates/pop"
    "github.com/markbates/validate"
    "github.com/satori/go.uuid"
)

type Like struct {
    ID        uuid.UUID `json:"id" db:"id"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
    UserID    uuid.UUID `json:"user_id" db:"user_id"`
    TweetID   uuid.UUID `json:"tweet_id" db:"tweet_id"`
}

// String is not required by pop and may be deleted
func (l Like) String() string {
    jl, _ := json.Marshal(l)
    return string(jl)
}

// Likes is not required by pop and may be deleted
type Likes []Like

// String is not required by pop and may be deleted
func (l Likes) String() string {
    jl, _ := json.Marshal(l)
    return string(jl)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *Like) Validate(tx *pop.Connection) (*validate.Errors, error) {
    return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *Like) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
    return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *Like) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
    return validate.NewErrors(), nil
}
