package view

import (
	"net/http"
)

type Error interface {
	PresentError(rw http.ResponseWriter,
		httpCode int,
		applicationCode int,
		message string)
}
