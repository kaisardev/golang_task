package create

import (
	"golang_test_task/helper"
	"golang_test_task/service"
	"net/http"
)

func Handler(response http.ResponseWriter, request *http.Request) {
	dto := &DTO{}

	if !helper.HandlePOST(response, request, dto) {
		return
	}

	if !helper.Validate(dto, response) {
		return
	}

	// We can add auth middleware to all routes where needed

	// working with some 3-rd party service
	task, err := service.Create()
	if err != nil {
		helper.RespondError(response, err, http.StatusInternalServerError)
		return
	}

	helper.RespondJSON(response, task)
}
