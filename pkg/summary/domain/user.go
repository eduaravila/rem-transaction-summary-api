package domain

import "errors"

type User struct {
	id    string
	name  string
	email string
}

func NewUser(id, name, email string) (*User, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	return &User{
		id:    id,
		name:  name,
		email: email,
	}, nil
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func DecodeUserFromCSV(record []string) (*User, error) {
	return NewUser(record[0], record[1], record[2])
}
