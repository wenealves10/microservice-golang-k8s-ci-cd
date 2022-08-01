package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/wenealves10/ddd-golang/domain/customer"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := customer.NewCustomer("John Doe")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cust,
		}}

	testCases := []testCase{
		{
			name:        "No Customer By ID",
			id:          uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		cust        string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Customer",
			cust:        "John Doe",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				customers: map[uuid.UUID]customer.Customer{},
			}

			cust, err := customer.NewCustomer(tc.cust)

			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(cust)

			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}

			if found.GetName() != cust.GetName() {
				t.Errorf("expected customer %v, got %v", cust.GetName(), found.GetName())
			}

		})
	}
}
