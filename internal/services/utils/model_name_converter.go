package utils

import (
	"github.com/gertd/go-pluralize"
	"strings"
)

func getSingularName(pluralWord string) string {
	pluralize := pluralize.NewClient()
	return pluralize.Singular(pluralWord)
}

func GetModelName(table string) string {
	modelName := getSingularName(table)
	modelName = strings.ToUpper(modelName[:1]) + modelName[1:]
	return modelName
}
