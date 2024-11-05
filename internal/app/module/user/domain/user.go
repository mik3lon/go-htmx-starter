package user_domain

import (
	"time"
)

type User struct {
	id             string
	username       string
	email          string
	hashedPassword string
	name           string
	surname        string
	role           string
	createdAt      time.Time
	updatedAt      time.Time
}

// CreateUser creates a new User entity.
func CreateUser(id, username, email, password, name, surname, role string) *User {
	return &User{
		id:             id,
		username:       username,
		email:          email,
		hashedPassword: password,
		name:           name,
		surname:        surname,
		role:           role,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
	}
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Email() string {
	return u.email
}

func (u *User) HashedPassword() string {
	return u.hashedPassword
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Surname() string {
	return u.surname
}

func (u *User) Role() string {
	return u.role
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func FromPrimitives(
	id,
	username,
	email,
	hashedPassword,
	name,
	surname,
	role string,
	createdAt, updatedAt time.Time,
) *User {
	return &User{
		id:             id,
		username:       username,
		email:          email,
		hashedPassword: hashedPassword,
		name:           name,
		surname:        surname,
		role:           role,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}
}
