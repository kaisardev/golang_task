package helper

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"io"
	"log"
	"net/http"
)

func RespondError(response http.ResponseWriter, err error, statusCode int) {
	errorMessage := err.Error()

	message := make(map[string]interface{})
	message["error"] = errorMessage

	if statusCode < 100 || statusCode >= 600 {
		response.WriteHeader(522)
		_, _ = response.Write([]byte("{\"error\": \"" + errorMessage + "\"}"))
		return
	}

	response.WriteHeader(statusCode)
	resp, err := json.Marshal(&message)
	if err != nil {
		log.Printf("helper errorResponse could not write error message\n")
		return
	}

	_, _ = response.Write(resp)
}

func RespondJSON(response http.ResponseWriter, body interface{}) {
	byteResp, err := json.Marshal(body)
	if err != nil {
		RespondError(response, err, http.StatusBadRequest)
		return
	}

	_, err = response.Write(byteResp)
	if err != nil {
		RespondError(response, err, http.StatusInternalServerError)
	}
}

func HandlePOST(response http.ResponseWriter, request *http.Request, dto interface{}) bool {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		RespondError(response, err, http.StatusBadRequest)
		return false
	}

	err = json.Unmarshal(body, &dto)
	if err != nil {
		RespondError(response, err, http.StatusBadRequest)
		return false
	}

	return true
}

func Validate(validatable validation.Validatable, response http.ResponseWriter) bool {
	err := validatable.Validate()
	if err != nil {
		RespondError(response, err, http.StatusUnprocessableEntity)
		return false
	}

	return true
}
