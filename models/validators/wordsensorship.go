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

func (holder *WordCensorship) IsValid(errors *validate.Errors) {
    var censortWord = "fuck"
    if strings.Contains(strings.TrimSpace(holder.Field), censortWord) {
        errors.Add(validators.GenerateKey(holder.Name), fmt.Sprintf("You shouldn't say '%s' ", censortWord))
    }
}
