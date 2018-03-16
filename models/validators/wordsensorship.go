package validators

import (
    "fmt"
    "strings"

    "github.com/gobuffalo/validate"
    "github.com/gobuffalo/validate/validators"
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
