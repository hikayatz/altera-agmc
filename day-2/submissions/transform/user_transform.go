package transform

import "github.com/hikayat13/alterra-agcm/day-2/submission/models"

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Photo string `json:"photo" form:"photo"`
}

func UserTransform(u *models.User) *User {
	return &User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Photo: u.Photo,
	}
}
