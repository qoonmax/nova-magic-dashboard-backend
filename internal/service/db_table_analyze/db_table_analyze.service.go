package db_table_analyze

import (
	v1 "magic-dashboard-backend/internal/controller/http/v1"

	"github.com/gertd/go-pluralize"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (a *Service) GetSingularFromPlural(pluralWord string) string {

	pluralize := pluralize.NewClient()

	return pluralize.Singular(pluralWord)
}

func (a *Service) GetTableNames(tables []v1.Table) []string {
	tableNames := make([]string, len(tables))
	for i, table := range tables {
		tableNames[i] = table.Name
	}
	return tableNames
}

func (a *Service) GetEntitiesNamesFromTableNames(tables []v1.Table) []string {
	entityNames := make([]string, len(tables))

	for i, table := range tables {
		entityNames[i] = a.GetSingularFromPlural(table.Name)
	}
	return entityNames
}

// word := "People"

// fmt.Printf("IsPlural(%s)   => %t\n", word, pluralize.IsPlural(word))
// fmt.Printf("IsSingular(%s) => %t\n", word, pluralize.IsSingular(word))
// fmt.Printf("Plural(%s)     => %s\n", word, pluralize.Plural(word))
// fmt.Printf("Singular(%s)   => %s\n", word, pluralize.Singular(word))
