package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type successConstant struct {
	OK Success
}

var SuccessConstant successConstant = successConstant{
	OK: Success{
		Response: successResponse{
			Meta: Meta{
				Status:  S_SUCCESS,
				Message: "Request successfully proceed",
				Code:    http.StatusOK,
			},
			Data: nil,
		},
	},
}

type Success struct {
	Response successResponse `json:"response"`
	Code     int             `json:"code"`
}

type successResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func SuccessBuilder(res *Success, data interface{}) *Success {
	res.Response.Data = data
	return res
}

func CustomSuccessBuilder(code int, data interface{}, message string) *Success {
	return &Success{
		Response: successResponse{
			Meta: Meta{
				Status:  S_SUCCESS,
				Message: message,
				Code:    code,
			},
			Data: data,
		},
	}
}

func SuccessResponse(data interface{}) *Success {
	return SuccessBuilder(&SuccessConstant.OK, data)
}

func (s *Success) Send(c echo.Context) error {
	return c.JSON(s.Code, s.Response)
}