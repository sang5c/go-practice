package method

type account struct {
	balance int
	firstName string
	lastName string
}

func (a *account) withdrawPointer(amount int)  {
	a.balance -= amount
}

func (a account) withdrawValue(amount int) {
	a.balance -= amount
}

func (a account) withdrawReturnValue(amount int) account  {
	a.balance -= amount
	return a
}
