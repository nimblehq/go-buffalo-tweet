package validators

import (
    "github.com/markbates/validate"
    "strings"
    "github.com/markbates/validate/validators"
    "fmt"
)

type WordCensorship struct {
    Name  string
    Field string
}

func (v *WordCensorship) IsValid(errors *validate.Errors) {
    if strings.Contains(strings.TrimSpace(v.Field), "fuck") {
        errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("You shouldn't say '%s' ", "fuck"))
    }
}