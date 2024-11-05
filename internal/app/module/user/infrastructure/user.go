package user_infrastructure

import (
	user_domain "go-boilerplate/internal/app/module/user/domain"
	"time"
)

type User struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	Username       string    `json:"username" gorm:"uniqueIndex;size:255"`
	Email          string    `json:"email" gorm:"uniqueIndex;size:255"`
	HashedPassword string    `json:"-" gorm:"size:255"`
	Name           string    `json:"name" gorm:"size:255"`
	Surname        string    `json:"surname" gorm:"size:255"`
	Role           string    `json:"role" gorm:"size:100"`
	CreatedAt      time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func FromUser(u *user_domain.User) *User {
	return &User{
		ID:             u.Id(),
		Username:       u.Username(),
		Email:          u.Email(),
		HashedPassword: u.HashedPassword(),
		Name:           u.Name(),
		Surname:        u.Surname(),
		Role:           u.Role(),
		CreatedAt:      u.CreatedAt(),
		UpdatedAt:      u.UpdatedAt(),
	}
}

func ToDomainUser(u User) *user_domain.User {
	return user_domain.FromPrimitives(
		u.ID,
		u.Username,
		u.Email,
		u.HashedPassword,
		u.Name,
		u.Surname,
		u.Role,
		u.CreatedAt,
		u.UpdatedAt)
}
