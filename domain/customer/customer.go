package customer

import (
	"errors"

	"github.com/google/uuid"

	"github.com/wenealves10/tavern"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

type Customer struct {
	person       *tavern.Person
	products     []*tavern.Item
	transactions []tavern.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &tavern.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*tavern.Item, 0),
		transactions: make([]tavern.Transaction, 0),
	}, nil

}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}

	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}

	c.person.Name = name
}
