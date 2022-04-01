package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"autoCreateTime"`
	UpdatedAt time.Time `json:"autoUpdateTime"`
}

type Followers struct {
	User_ID     uint64    `json:"user_id,omitempty"`
	Follower_ID uint64    `json:"follower_id,omitempty"`
	CreatedAt   time.Time `json:"autoCreateTime"`
	UpdatedAt   time.Time `json:"autoUpdateTime"`
}

func (user *User) Prepare(procedure string) error {
	if err := user.validate(procedure); err != nil {
		return err
	}

	if err := user.format(procedure); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(procedure string) error {
	if user.Name == "" {
		return errors.New("name must not be empty")
	}

	if user.Username == "" {
		return errors.New("username must not be empty")
	}

	if user.Email == "" {
		return errors.New("email must not be empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email not valid")
	}

	if procedure == "register" && user.Password == "" {
		return errors.New("password must not be empty")
	}

	return nil
}

func (user *User) format(procedure string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	if procedure == "register" {
		hashPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPassword)
	}

	return nil
}
