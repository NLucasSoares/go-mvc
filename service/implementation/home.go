package implementation

import (
	"test/model"
	"test/model/repository"
)

type Home struct {
	Repository repository.Home
}

func NewHome(repository repository.Home) *Home {
	return &Home{
		Repository: repository,
	}
}

func (home *Home) GetHome(name string) (*model.Home, error) {
	return home.Repository.FindHome(name)
}
