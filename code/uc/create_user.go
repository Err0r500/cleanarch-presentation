package uc

import (
	"errors"
	"time"
)

type User struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

type Interactor struct {
	UserRW UserReadWriter
}

type UserReadWriter interface {
	Create(User) (id int, err error)
	Save(User) error
	GetByID(id string) (User, error)
}

func (i Interactor) CreateUser(name string, dateOfBirth time.Time) error {
	user := User{
		Name:        name,
		DateOfBirth: dateOfBirth,
	}

	if err := user.Check(); err != nil {
		return err
	}

	id, err := i.UserRW.Create(user)
	if err != nil {
		return err
	}
	if id == 0 {
		return errors.New("user creation failed without raising an error")
	}

	return nil
}

func (user User) Check() error {
	return nil
}
