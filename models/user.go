package models

type User struct {
	ID     int            `json:"id" form:"id" gorm:"primaryKey"`
	Name   string         `json:"name" form:"name" gorm:"not null"`
	Locker LockerResponse `json:"locker"` //one to one
	Posts  []PostResponse `json:"posts" ` //one to many
}

type UserResponse struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func (UserResponse) TableName() string {
	return "users"
}
