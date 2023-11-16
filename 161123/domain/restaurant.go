package domain



type restaurant struct {
	name     string
	capacity int64
	price	int64
}

func CreateRestaurant(name string, capacity, price int64) *restaurant {
	return &restaurant{
		name:     name, 
		capacity: capacity,
		price:      price,
	}
}

func (r *restaurant) GetName() string {
	return r.name
}
func (r *restaurant) GetCapacity() int64 {
	return r.capacity
}
func (r *restaurant) GetPrice() int64 {
	return r.price
}

// hello 2 test