package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wenealves10/microservice-golang-k8s-ci-cd/aggregate"
)

var (
	ErrProductNotFound      = errors.New("the product was not found")
	ErrProductAlreadyExists = errors.New("the product already exists")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(aggregate.Product) error
	Update(aggregate.Product) error
	Delete(id uuid.UUID) error
}
