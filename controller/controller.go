package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"test/service"
	"test/view"
	"test/view/implementation"
)

type Controller struct {
	router *mux.Router

	errorView view.Error

	homeService service.Home
	homeView    view.Home
}

func New(homeService service.Home,
	homeView view.Home) *Controller {
	controller := &Controller{
		homeService: homeService,
		homeView:    homeView,

		errorView: implementation.NewError(),
	}

	controller.createRoute()

	return controller
}

func (controller *Controller) createRoute() {
	controller.router = mux.NewRouter()

	controller.router.HandleFunc("/home/{name}", controller.HomeHandler)

	go func() {
		_ = http.ListenAndServe(":8080",
			controller.router)
	}()
}

func (controller *Controller) HomeHandler(rw http.ResponseWriter,
	request *http.Request) {
	if name, ok := mux.Vars(request)["name"]; ok {
		if homeModel, err := controller.homeService.GetHome(name); err == nil {
			controller.homeView.PresentHome(rw,
				homeModel)
		} else {
			controller.errorView.PresentError(rw,
				404,
				1234,
				"Home not found: "+err.Error())
		}
	} else {
		controller.errorView.PresentError(rw,
			400,
			5789,
			"Home name parameter not found!")
	}
}
