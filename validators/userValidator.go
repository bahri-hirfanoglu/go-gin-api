package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidatePictureUrl(field validator.FieldLevel) bool {
	str := fmt.Sprintf("%s", field.Field())
	format := str[len(str)-3 : len(str)]
	return strings.Compare(strings.TrimSpace(format), "png") == 0
}
