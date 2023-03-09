package get

import (
	"github.com/gorilla/mux"
	"golang_test_task/helper"
	"golang_test_task/service"
	"net/http"
)

const (
	paramID = "id"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params[paramID]

	if err := validate(); err != nil {
		helper.RespondError(response, err, http.StatusBadRequest)
		return
	}

	// Going to some 3-rd party service
	task, err := service.Get(id)
	if err != nil {
		helper.RespondError(response, err, http.StatusInternalServerError)
		return
	}

	helper.RespondJSON(response, task)
}

func validate() error {
	// Must be validation depends on ID type

	return nil
}
