package manipulator

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/theseceng/manipulator/utils"
)

func HandleUpper(value string) string {
	return strings.ToUpper(value)
}

func HandleLower(value string) string {
	return strings.ToLower(value)
}

func HandleSnake(value string) string {
	return strcase.ToSnake(value)
}

func HandleCamel(value string) string {
	return strcase.ToCamel(value)
}

func HandleLowerCamel(value string) string {
	return strcase.ToLowerCamel(value)
}

func HandleTitleCase(value string, rules string) string {
	return utils.ToTitleCase(value, rules)
}
