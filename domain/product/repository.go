package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("the product was not found")
	ErrProductAlreadyExists = errors.New("the product already exists")
)

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (Product, error)
	Add(Product) error
	Update(Product) error
	Delete(id uuid.UUID) error
}
