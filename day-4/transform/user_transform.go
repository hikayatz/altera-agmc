package transform

import "github.com/hikayat13/alterra-agcm/day-2/submission/models"

type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Photo     string `json:"photo" form:"photo"`
	CreatedBy int    `json:"created_by" form:"photo"`
	Token     string `json:"token,omitempty"`
}

func UserTransform(u *models.User) *User {
	return &User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Photo:     u.Photo,
		CreatedBy: u.CreatedBy,
	}
}
