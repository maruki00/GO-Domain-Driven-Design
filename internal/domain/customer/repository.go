package Customer

import "errors"

var (
	CustomerNotFoundErr    = errors.New("Customer Not found")
	FailedToAddCustomerErr = errors.New("Failed to add Customer")
	UpdateCustomerErr      = errors.New("Could not update to custmer")
)

type CustomerRespository interface {
	Get(int) (aggrigate.Customer, error)
	Add(aggrigate.Customer) error
	Update(aggrigate.Customer) error
}
