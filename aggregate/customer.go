package aggregate

import (
	"github.com/wenealves10/microservice-golang-k8s-ci-cd/entity"
	"github.com/wenealves10/microservice-golang-k8s-ci-cd/valueobject"
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []*valueobject.Transaction
}
