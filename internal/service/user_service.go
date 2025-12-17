package service

import (
	"context"
	"time"

	"github.com/gitams17/ainyx-assignment/db/sqlc"
	"github.com/gitams17/ainyx-assignment/internal/models"
	"github.com/gitams17/ainyx-assignment/internal/repository"
	"github.com/jackc/pgx/v5/pgtype" // Required for pgtype.Date
)

// Fix: ListUsers should return a slice []models.UserResponse, not a single struct
type UserService interface {
	CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error)
	GetUser(ctx context.Context, id int64) (*models.UserResponse, error)
	ListUsers(ctx context.Context, page int32, limit int32) ([]models.UserResponse, error)
	UpdateUser(ctx context.Context, id int64, req models.UpdateUserRequest) (*models.UserResponse, error)
	DeleteUser(ctx context.Context, id int64) error
}

type userService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{repo: repo}
}

func CalculateAge(dob time.Time, now time.Time) int {
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

// Helper to convert time.Time to pgtype.Date
func toPgDate(t time.Time) pgtype.Date {
	return pgtype.Date{Time: t, Valid: true}
}

func (s *userService) CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return nil, err
	}

	// Fix: Use helper to convert time.Time to pgtype.Date
	user, err := s.repo.CreateUser(ctx, db.CreateUserParams{
		Name: req.Name,
		Dob:  toPgDate(dob),
	})
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Time.Format("2006-01-02"), // Fix: Access .Time field
		Age:  CalculateAge(user.Dob.Time, time.Now()),
	}, nil
}

func (s *userService) GetUser(ctx context.Context, id int64) (*models.UserResponse, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Time.Format("2006-01-02"),
		Age:  CalculateAge(user.Dob.Time, time.Now()),
	}, nil
}

// Fix: Return slice []models.UserResponse
func (s *userService) ListUsers(ctx context.Context, page int32, limit int32) ([]models.UserResponse, error) {
	offset := (page - 1) * limit
	users, err := s.repo.ListUsers(ctx, db.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	// Fix: Initialize slice
	var response []models.UserResponse
	for _, u := range users {
		response = append(response, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob.Time.Format("2006-01-02"),
			Age:  CalculateAge(u.Dob.Time, time.Now()),
		})
	}
	return response, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int64, req models.UpdateUserRequest) (*models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil { return nil, err }

	user, err := s.repo.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  toPgDate(dob),
	})
	if err != nil { return nil, err }

	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Time.Format("2006-01-02"),
		Age:  CalculateAge(user.Dob.Time, time.Now()),
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}