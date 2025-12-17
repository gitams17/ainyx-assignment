package repository

import (
	"context"
	"github.com/gitams17/ainyx-assignment/db/sqlc" // Imported generated code
)

// Repository defines the interface for database operations
// This interface makes the Service layer testable via mocking
type Repository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUser(ctx context.Context, id int64) (db.User, error)
	ListUsers(ctx context.Context, arg db.ListUsersParams) (db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
	DeleteUser(ctx context.Context, id int64) error
	CountUsers(ctx context.Context) (int64, error)
}