package validators

import (
    "testing"
    "github.com/gobuffalo/suite"
    "github.com/markbates/validate"
)

type ModelSuite struct {
    *suite.Model
}

// TODO: fix this
func (ms *ModelSuite) Test_WordCensorship_Validator(t *testing.T) {
    censor := WordCensorship{
        Name: "name",
        Field: "fuck",
    }
    errs := validate.NewErrors()
    before := errs.Count()
    censor.IsValid(errs)
    after := errs.Count()
    ms.Equal(before, after)
}