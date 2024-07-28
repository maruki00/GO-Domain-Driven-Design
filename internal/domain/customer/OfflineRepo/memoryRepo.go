package offlinerepo

type CustomerMemoryRepo struct {
	items map[int]aggrigate.Custmer
	sync.Mutex
}

//Get(int) (aggrigate.Customer, error)
//Add(aggrigate.Customer) error
//Update(aggrigate.Customer) error

func New() *CustomerMemoryRepo {
	return &CustomerMemoryRepo{
		items : make(map[int]aggrigate.Custmer, 0)
	}
}

func (o *CustomerMemoryRepo)Get(id int) (*aggrigate.Customer, error) {

	item,ok := o.items[id]
	if ok != nil {
		return nil,errors.New("Could not get the customer")
	}

	return item, nil

}

func (o *CustomerMemoryRepo)Add(customer aggrigate.Customer) error {


}

func (o *CustomerMemoryRepo)Update(customer aggrigate.Customer) error {


}