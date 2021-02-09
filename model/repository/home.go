package repository

import (
	"test/model"
)

type Home interface {
	FindHome(name string) (*model.Home, error)
}
