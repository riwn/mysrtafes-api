package errors

import (
	"encoding/json"
	"mysrtafes-backend/pkg/errors"
	"net/http"
)

type response struct {
	Error interface{} `json:"error"`
}

type errorBase struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

type errorWithInformation struct {
	errorBase
	Info          errorInformation `json:"info"`
	InvalidParams interface{}      `json:"invalid-params"`
}

type errorInformation struct {
	ErrorCode    errors.Code    `json:"error_code"`
	ErrorMessage errors.Message `json:"error_message"`
}

func WriteError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case errors.InvalidRequestError,
		errors.InvalidValidateError:
		writeError(w, http.StatusBadRequest, err)
	case errors.NotFoundError:
		writeError(w, http.StatusNotFound, err)
	case errors.ForbiddenError:
		writeError(w, http.StatusForbidden, err)
	case errors.UnauthorizedError:
		writeError(w, http.StatusUnauthorized, err)
	case errors.UnsupportedMediaTypeError:
		writeError(w, http.StatusUnsupportedMediaType, err)
	case errors.InternalServerErrorError:
		writeError(w, http.StatusInternalServerError, err)
	default:
		writeError(w, http.StatusInternalServerError, err)
	}
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	body := response{}
	switch err := err.(type) {
	case errors.Informator:
		var code errors.Code
		var message errors.Message
		var invalidParams interface{}
		if err.Information() != nil {
			code, message = err.Information().Code.Detail()
			invalidParams = err.Information().Problem
		}
		body.Error = errorWithInformation{
			errorBase: errorBase{
				Title:  http.StatusText(statusCode),
				Status: statusCode,
				Detail: err.Error(),
			},
			Info: errorInformation{
				ErrorCode:    code,
				ErrorMessage: message,
			},
			InvalidParams: invalidParams,
		}
	default:
		body.Error = errorBase{
			Title:  http.StatusText(statusCode),
			Status: statusCode,
			Detail: err.Error(),
		}
	}

	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&body)
}
