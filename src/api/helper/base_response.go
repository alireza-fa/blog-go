package helper

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ValidationError struct {
	Key   string `json:"key"`
	Error string `json:"error"`
}

type BaseHttpResponse struct {
	Success          bool              `json:"success"`
	Result           interface{}       `json:"result"`
	StatusCode       int               `json:"statusCode"`
	Error            interface{}       `json:"error,omitempty"`
	ValidationErrors []ValidationError `json:"validationErrors,omitempty"`
}

func BaseResponse(w http.ResponseWriter, result interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	response := BaseHttpResponse{
		Success:    true,
		Result:     result,
		StatusCode: statusCode,
	}

	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		BaseResponseWithError(w, nil, http.StatusInternalServerError, err)
		panic(err)
	}
}

func BaseResponseWithError(w http.ResponseWriter, result interface{}, statusCode int, errorInfo error) {
	w.Header().Set("Content-Type", "application/json")

	response := BaseHttpResponse{
		Success:    false,
		Result:     result,
		StatusCode: statusCode,
		Error:      errorInfo.Error(),
	}

	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		BaseResponseWithError(w, nil, http.StatusInternalServerError, err)
		return
	}
}

func BaseResponseWithValidationError(w http.ResponseWriter, result interface{}, statusCode int, errorInfo error) {
	w.Header().Set("Content-Type", "application/json")

	response := BaseHttpResponse{
		Success:    false,
		Result:     result,
		StatusCode: statusCode,
	}

	errors := strings.Split(errorInfo.Error(), "\n")

	for _, item := range errors {
		splitError := strings.Split(item, "Error:")

		key := strings.Replace(splitError[0], "Key: ", "", 1)
		key = strings.TrimSpace(key)
		key = strings.Trim(key, "'")

		response.ValidationErrors = append(response.ValidationErrors, ValidationError{
			Key:   key,
			Error: splitError[1],
		})
	}

	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		BaseResponseWithError(w, nil, http.StatusInternalServerError, err)
		return
	}
}
