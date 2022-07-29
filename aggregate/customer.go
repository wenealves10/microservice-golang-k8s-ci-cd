package aggregate

import (
	"errors"

	"github.com/google/uuid"

	"github.com/wenealves10/microservice-golang-k8s-ci-cd/entity"
	"github.com/wenealves10/microservice-golang-k8s-ci-cd/valueobject"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil

}
