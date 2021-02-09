package implementation

import (
	"test/model"
	"test/model/errors"
)

type Home struct {
	// let's suppose home data are fetched from db for example. for now I use data
}

func NewHome() *Home {
	return &Home{}
}

var (
	// where the key is the name
	data = map[string]*model.Home{
		"first": {
			Name:    "first",
			Address: "44 bis rue de puiseux 95490 VAUREAL",
		},
		"second": {
			Name:    "second",
			Address: "128 bd de l'oise 95490 VAUREAL",
		},
	}
)

func (home *Home) FindHome(name string) (*model.Home, error) {
	if foundHome, ok := data[name]; ok {
		return foundHome, nil
	} else {
		return nil, errors.NotFoundError
	}
}
