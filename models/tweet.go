package models

import (
    "encoding/json"
    "time"

    validators2 "github.com/bufftwitt/models/validators"
    "github.com/gobuffalo/pop"
    "github.com/gobuffalo/validate"
    "github.com/gobuffalo/validate/validators"
    "github.com/gobuffalo/uuid"
)

type Tweet struct {
    ID        uuid.UUID `json:"id" db:"id"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
    UserID    uuid.UUID `json:"user_id" db:"user_id"`
    Message   string    `json:"message" db:"message"`
    Author    string    `json:"-" db:"author_name" form:"-" select:"(select name from users where id = tweets.user_id) as author_name" rw:"r"`
    LikeCount int       `json:"-" db:"like_count" form:"-" select:"(select count(*) from likes where tweet_id = tweets.id) as like_count" rw:"r"`
}

// String is not required by pop and may be deleted
func (t Tweet) String() string {
    jt, _ := json.Marshal(t)
    return string(jt)
}

// Tweets is not required by pop and may be deleted
type Tweets []Tweet

// String is not required by pop and may be deleted
func (t Tweets) String() string {
    jt, _ := json.Marshal(t)
    return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Tweet) Validate(tx *pop.Connection) (*validate.Errors, error) {
    isEmpty := validate.Validate(
        &validators.StringIsPresent{Field: t.Message, Name: "Message"},
    )
    if isEmpty != nil {
        return validate.Validate(
            &validators2.WordCensorship{Field: t.Message, Name: "Message"},
        ), nil
    }
    return isEmpty, nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Tweet) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
    return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Tweet) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
    return validate.NewErrors(), nil
}
