package helper

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}
type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) *response {
	return &response{
		Meta: meta{
			Message: message,
			Code:    code,
			Status:  status,
		},
		Data: data,
	}
}
