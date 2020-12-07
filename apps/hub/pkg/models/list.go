package models

import (
	"github.com/google/uuid"
)

type List struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Owner uuid.UUID `json:"owner"`
}
