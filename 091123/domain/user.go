package domain

// hi

type user struct {
	name     string
	lastName string
	age      int64
	balance  int64
}

func CreateUser(name, lastName string, age, balance int64) *user {
	return &user{
		name:     name,
		lastName: lastName,
		age:      age,
		balance:  balance,
	}
}

func (u *user) GetName() string {
	return u.name
}
func (u *user) GetLastName() string {
	return u.lastName
}
func (u *user) GetAge() int64 {
	return u.age
}
func (u *user) GetBalance() int64 {
	return u.balance
}

func (u *user) AddMoney(amount int64) {
	u.balance += amount
}

func (u *user) SetName(name string) {
	u.name = name
}
func (u *user) SetLastName(lastName string) {
	u.lastName = lastName
}
func (u *user) SetAge(age int64) {
	u.age = age
}
func (u *user) SetBalance(balance int64) {
	u.balance = balance
}