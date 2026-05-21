package models

type UserID int
type UserName = string

type User struct {
	ID       UserID `structtomap:"id"`
	Username string `structtomap:"username"`
	Email    string `structtomap:"email"`
	Active   bool   `structtomap:"active"`
}

type Users []User
