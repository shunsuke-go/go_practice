package entity

import (
	"fmt"
	"time"
)

	
type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Createdat time.Time  `json:"-"`
	Updatedat time.Time  `json:"-"`
}

func (u *User) AddPrefix() string {
	return fmt.Sprintf("%sさん", u.Name)
}