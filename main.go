package main

import (
	"test/controller"
	repositoryImplementation "test/model/repository/implementation"
	serviceImplementation "test/service/implementation"
	viewImplementation "test/view/implementation"
	"time"
)

func main() {
	_ = controller.New(
		serviceImplementation.NewHome(
			repositoryImplementation.NewHome(),
		),
		viewImplementation.NewHome(),
	)

	for {
		time.Sleep(time.Second)
	}
}
