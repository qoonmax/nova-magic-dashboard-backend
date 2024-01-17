package utils

import "github.com/gertd/go-pluralize"

func GetSingularName(pluralWord string) string {
	pluralize := pluralize.NewClient()
	return pluralize.Singular(pluralWord)
}
