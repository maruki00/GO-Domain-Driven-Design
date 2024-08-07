package services

type OrderConfiguration func(os *orderService) error

type OrderService struct {
	customers customer.CustomerRespository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {

	os := OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr custmer.CustomerRespository) OrderConfiguration {
	return func(os *orderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository(cr custmer.CustomerRespository) OrderConfiguration {
	cr := offlinerepo.New()
	return WithCustomerRepository(cr)
}

func (o *OrderService) CreateOrder(customerId int, productsId []int) error {
	c, err := o.customers
}
