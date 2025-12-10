package product

import (
	"time"

	"github.com/rs/xid"
)

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewProduct(name, description string) (*Product, error) {
	newId := xid.New().String()

	return &Product{
		ID:          newId,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}, nil
}
