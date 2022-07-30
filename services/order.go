package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/wenealves10/microservice-golang-k8s-ci-cd/aggregate"
	"github.com/wenealves10/microservice-golang-k8s-ci-cd/domain/customer"
	"github.com/wenealves10/microservice-golang-k8s-ci-cd/domain/customer/memory"
	"github.com/wenealves10/microservice-golang-k8s-ci-cd/domain/product"
	prodmemory "github.com/wenealves10/microservice-golang-k8s-ci-cd/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOrderService(cgfs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cgfs {

		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()

	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		pr := prodmemory.New()

		// Add Items to repo
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// Get the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	// Get each Product, Ouchie, We need a ProductRepository
	var products []aggregate.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}

	// All Products exists in store, now we can create the order
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return price, nil
}
