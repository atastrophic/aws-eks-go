package repositories

import (
	"github.com/atastrophic/go-ms-with-eks/pkg/exception"
	"github.com/atastrophic/go-ms-with-eks/pkg/source"
	"github.com/google/uuid"
)

type ListRepository struct {
	source source.Source
}

func NewListRepository(source source.Source) *ListRepository {
	return &ListRepository{
		source: source,
	}
}

func (r *ListRepository) Create(name string, owner uuid.UUID) uuid.UUID {

	id, err := uuid.NewUUID()
	exception.WithMessageAndError("unable to generate UUID", err)
	r.source.Exec("create-list", id, name, owner)
	return id
}
