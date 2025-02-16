package entity

import (
	"time"

	"github.com/akira-saneyoshi/task-app/domain/object/value"
)

type User struct {
	ID        *value.ID       `json:"id"`
	Name      *value.Name     `json:"name"`
	Email     *value.Email    `json:"email"`
	Password  *value.Password `json:"password"`
	IsActive  bool            `json:"is_active"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func (u *User) Validate() error {
	if err := u.ID.Validate(); err != nil {
		return err
	}
	if err := u.Name.Validate(); err != nil {
		return err
	}
	if err := u.Email.Validate(); err != nil {
		return err
	}
	if err := u.Password.Validate(); err != nil {
		return err
	}
	return nil
}
