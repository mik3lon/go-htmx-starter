package user_infrastructure

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	user_domain "go-boilerplate/internal/app/module/user/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresUserRepository is a Postgres implementation of UserRepository using Gorm.
type PostgresUserRepository struct {
	DB *gorm.DB
}

// NewPostgresUserRepository initializes a new Postgres user repository.
func NewPostgresUserRepository(dsn string) (*PostgresUserRepository, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Ensure the User table exists
	if err = db.AutoMigrate(&user_domain.User{}); err != nil {
		return nil, err
	}

	return &PostgresUserRepository{
		DB: db,
	}, nil
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *user_domain.User) error {
	result := r.DB.WithContext(ctx).Save(user)
	if result.Error != nil {
		// Check if the error is a unique constraint violation
		var pgErr *pgconn.PgError
		errors.As(result.Error, &pgErr)
		if pgErr.Code == "23505" {
			return user_domain.NewUserAlreadyExists(user.Email)
		}
		return fmt.Errorf("failed to save user: %w", result.Error)
	}
	return nil
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*user_domain.User, error) {
	var user User
	result := r.DB.First(&user, "email = ?", email)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, user_domain.ErrUserNotFound
	}
	return ToDomainUser(user), result.Error
}

func (r *PostgresUserRepository) FindAll(ctx context.Context) (user_domain.UserList, error) {
	var users []User
	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	userList := make(user_domain.UserList, len(users))
	for i, u := range users {
		userList[i] = ToDomainUser(u)
	}
	return userList, nil
}

func isUniqueViolation(err error) bool {
	var pgError *pgconn.PgError
	if errors.As(err, &pgError) {
		return pgError.Code == "23505"
	}
	return false
}
