package request

type (
	UserCreate struct {
		Name     string `json:"name"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
		UserId   int    `json:",omitempty"`
	}
)
