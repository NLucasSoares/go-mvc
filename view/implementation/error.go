package implementation

import (
	"encoding/json"
	"net/http"
)

type errorPresentation struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Error struct{}

func NewError() *Error {
	return &Error{}
}

func (error *Error) PresentError(rw http.ResponseWriter,
	httpCode int,
	applicationCode int,
	message string) {
	errorPresentation := &errorPresentation{
		Code:    applicationCode,
		Message: "VIEW DOING ITS JOB: " + message,
	}
	view, _ := json.MarshalIndent(errorPresentation,
		"",
		"\t")
	rw.WriteHeader(httpCode)
	_, _ = rw.Write(view)
}
