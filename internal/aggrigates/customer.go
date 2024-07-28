package aggrigate


var (
	InvalidPersonError = errors.New("a customer has to have a valid name")
)
import (
	"Go-Domain-Driven-Design/internal/entity"
	"Go-Domain-Driven-Design/internal/valueobject"
)

type Customer struct {
	Person       *entity.Person
	produts       []*entity.Item
	Transactions []valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return &Customer{},InvalidPersonError
	}

	person := &entity.Person{
		Name : name,
		ID: 1
	}
	return &Customer{
		person: person,
		produts: make(entity.Item,0),
		Transactions: make(valueobject.Transaction, 0)
	}, 0
}
