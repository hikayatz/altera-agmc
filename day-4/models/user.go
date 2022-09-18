package models

type User struct {
	Common
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password,omitempty" form:"password"`
	Photo     string `json:"photo" form:"photo"`
	CreatedBy int    `json:"created_by"`
}

// TableName this is default and customize
func (u *User) TableName() string {
	return "users"
}
