package response

import "gorm.io/gorm"

type response struct {
	Meta Meta        `json:"Meta"`
	Data interface{} `json:"data,omitempty"`
}
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

const (
	S_SUCCESS         = "success"
	S_INVALID_REQUEST = "invalid_request"
	S_ERROR           = "error"
)

var S_NOT_FOUND = gorm.ErrRecordNotFound

func APIResponse(message string, code int, status string, data interface{}) *response {
	return &response{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  status,
		},
		Data: data,
	}
}