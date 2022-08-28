package model

import "time"

type User struct {
	UserId    string    `json:"-" db:"user_id,omitempty"`
	FullName  string    `json:"fullName,omitempty" db:"full_name,omitempty"`
	Email     string    `json:"email,omitempty" db:"email,omitempty"`
	Password  string    `json:"password,omitempty" db:"password,omitempty"`
	Role      string    `json:"-" db:"role,omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at,omitempty"`
	UpdateAt  time.Time `json:"-" db:"updated_at,omitempty"`
	Token     string    `json:"token,omitempty"`
}

func New() User {
	return User{}
}

func (u User) GetEmail() string {
	return u.Email
}
