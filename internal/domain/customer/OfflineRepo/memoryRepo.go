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

	if o.items == nil{
		o.Lock()
		o.items = make(map[int]aggrigate.Custmer, 0)
		o.Unlock()
	}

	if _, ok = c.items[custmer.ID]; ok {
		return errors.New("customer Already Exists")
	}
	o.Lock()
	o.items[c.ID] = c
	o.Unlock()
	return nil
}

func (o *CustomerMemoryRepo)Update(customer aggrigate.Customer) error {
	if _, ok = c.items[custmer.ID]; !ok {
		return errors.New("customer Doesnt Exists")
	}
	o.Lock()
	o.items[c.ID] = c
	o.Unlock()
	return nil

}