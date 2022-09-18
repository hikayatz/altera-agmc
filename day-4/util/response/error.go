package response

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	E_DUPLICATE            = "duplicate"
	E_NOT_FOUND            = "not_found"
	E_UNPROCESSABLE_ENTITY = "unprocessable_entity"
	E_UNAUTHORIZED         = "unauthorized"
	E_SERVER_ERROR         = "server_error"
	E_REQUEST_NOT_VALID    = "request_not_valid"
)

type Error struct {
	Code         int
	Response     errorResponse
	ErrorMessage error
}

type errorResponse struct {
	Meta  Meta   `json:"Meta"`
	Error string `json:"error,omitempty"`
}

type errorConstant struct {
	Duplicate                errorResponse
	NotFound                 errorResponse
	RouteNotFound            errorResponse
	UnprocessableEntity      errorResponse
	Unauthorized             errorResponse
	BadRequest               errorResponse
	Validation               errorResponse
	InternalServerError      errorResponse
	EmailOrPasswordIncorrect errorResponse
	ConvertionNotValid       errorResponse
	NotEnoughStock           errorResponse
}

var ErrorConstant errorConstant = errorConstant{
	Duplicate: errorResponse{
		Meta: Meta{
			Status:  S_INVALID_REQUEST,
			Message: "Created value already exists",
			Code:    http.StatusConflict,
		},
		Error: E_DUPLICATE,
	},

	EmailOrPasswordIncorrect: errorResponse{
		Meta: Meta{
			Status:  S_INVALID_REQUEST,
			Message: "Email or password is incorrect",
			Code:    http.StatusBadRequest,
		},
	},
	NotFound: errorResponse{
		Meta: Meta{
			Status:  S_INVALID_REQUEST,
			Message: "Data not found",
			Code:    http.StatusNotFound,
		},
		Error: E_NOT_FOUND,
	},
	RouteNotFound: errorResponse{
		Meta: Meta{
			Status:  S_INVALID_REQUEST,
			Message: "Route not found",
			Code:    http.StatusNotFound,
		},
		Error: E_NOT_FOUND,
	},
	UnprocessableEntity: errorResponse{
		Meta: Meta{
			Status:  S_INVALID_REQUEST,
			Message: "Invalid parameters or payload",
			Code:    http.StatusUnprocessableEntity,
		},
		Error: E_UNPROCESSABLE_ENTITY,
	},

	Unauthorized: errorResponse{
		Meta: Meta{
			Status:  S_INVALID_REQUEST,
			Message: "Unauthorized, please login or use different role",
			Code:    http.StatusUnauthorized,
		},
		Error: E_UNAUTHORIZED,
	},
	BadRequest: errorResponse{
		Meta: Meta{
			Status:  S_INVALID_REQUEST,
			Message: "Bad Request",
			Code:    http.StatusBadRequest,
		},
	},
	Validation: errorResponse{
		Meta: Meta{
			Status:  S_INVALID_REQUEST,
			Message: "Invalid parameters or payload",
			Code:    http.StatusBadRequest,
		},
	},
	InternalServerError: errorResponse{
		Meta: Meta{
			Status:  S_INVALID_REQUEST,
			Message: "Something bad happened",
			Code:    http.StatusInternalServerError,
		},
		Error: E_SERVER_ERROR,
	},
	ConvertionNotValid: errorResponse{
		Meta: Meta{
			Status:  S_ERROR,
			Message: "Invalid request",
			Code:    http.StatusBadRequest,
		},
		Error: E_REQUEST_NOT_VALID,
	},
}

func ErrorBuilder(res *errorResponse, message error) *Error {
	return &Error{
		ErrorMessage: message,
		Response:     *res,
	}
}

func (e *Error) Send(c echo.Context) error {
	return c.JSON(e.Response.Meta.Code, e.Response)
}

func ErrorResponse(err error) *Error {
	re, ok := err.(*Error)
	if ok {
		return re
	} else {
		return ErrorBuilder(&ErrorConstant.InternalServerError, err)
	}
}
func (e *Error) ParseToError() error {
	return e.ErrorMessage
}
func CustomErrorBuilder(code int, err string, message string) *Error {
	return &Error{
		Response: errorResponse{
			Meta: Meta{
				Status:  err,
				Message: message,
				Code:    code,
			},
		},
		Code: code,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code %d", e.Code)
}