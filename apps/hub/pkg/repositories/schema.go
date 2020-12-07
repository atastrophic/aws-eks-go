package repositories

import (
	"github.com/atastrophic/go-ms-with-eks/pkg/source"
)

type Schema struct {
	source source.Source
}

func NewSchema(source source.Source) *Schema {
	return &Schema{source: source}
}

func (s *Schema) Execute() {
	s.source.Exec("create-users-table")
	s.source.Exec("create-lists-table")
	s.source.Exec("create-todos-table")
}
