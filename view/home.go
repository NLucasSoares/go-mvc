package view

import (
	"net/http"
	"test/model"
)

type Home interface {
	PresentHome(rw http.ResponseWriter,
		home *model.Home)
}
