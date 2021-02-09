package implementation

import (
	"encoding/json"
	"net/http"
	"test/model"
)

type homePresentation struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Home struct{}

func NewHome() *Home {
	return &Home{}
}

func (home *Home) PresentHome(rw http.ResponseWriter,
	homeModel *model.Home) {
	homePresentation := &homePresentation{
		Name:    homeModel.Name,
		Address: homeModel.Address,
	}
	view, _ := json.MarshalIndent(homePresentation,
		"",
		"\t")
	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(view)
}
