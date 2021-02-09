package service

import (
	"test/model"
)

type Home interface {
	GetHome(name string) (*model.Home, error)
}
