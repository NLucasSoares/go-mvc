package model

type Home struct {
	Name string

	Address string
}

func NewHome(name string,
	address string) *Home {
	return &Home{
		Name: name,

		Address: address,
	}
}
