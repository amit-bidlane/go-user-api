package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	db "go-user-api/db/sqlc"
	"go-user-api/internal/models"
	"go-user-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func CalculateAge(dob time.Time) int {
	if dob.After(time.Now()) {
    	return 0
	}	
	today := time.Now()
	age := today.Year() - dob.Year()
	if today.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

func validateDOB(dobStr string) error {
	if strings.HasPrefix(dobStr, " ") || strings.HasSuffix(dobStr, " ") {
		return fmt.Errorf("date of birth cannot have leading or trailing spaces")
	}

	parts := strings.Split(dobStr, "-")
	if len(parts) != 3 {
		return fmt.Errorf("invalid date format, use YYYY-MM-DD")
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid month value")
	}

	day, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("invalid day value")
	}

	if month > 12 {
		return fmt.Errorf("month cannot be greater than 12")
	}

	if day > 31 {
		return fmt.Errorf("day cannot be greater than 31")
	}

	thirtyDayMonths := map[int]string{
		4: "April", 6: "June", 9: "September", 11: "November",
	}

	if name, ok := thirtyDayMonths[month]; ok && day == 31 {
		return fmt.Errorf("%s only has 30 days", name)
	}

	if month == 2 && day > 29 {
		return fmt.Errorf("February cannot have more than 28 or 29 days")
	}

	return nil
}

func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	if err := validateDOB(req.Dob); err != nil {
		return models.UserResponse{}, err
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("invalid date: day or month value does not exist")
	}

	if dob.After(time.Now()) {
		return models.UserResponse{}, fmt.Errorf("date of birth cannot be in the future")
	}

	user, err := s.repo.CreateUser(ctx, db.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
	}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (models.UserWithAgeResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return models.UserWithAgeResponse{}, err
	}

	return models.UserWithAgeResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  CalculateAge(user.Dob),
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	if err := validateDOB(req.Dob); err != nil {
		return models.UserResponse{}, err
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("invalid date: day or month value does not exist")
	}

	if dob.After(time.Now()) {
		return models.UserResponse{}, fmt.Errorf("date of birth cannot be in the future")
	}

	user, err := s.repo.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context, page, limit int) (models.PaginatedUsersResponse, error) {
	offset := (page - 1) * limit

	users, err := s.repo.ListUsers(ctx, db.ListUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return models.PaginatedUsersResponse{}, err
	}

	total, err := s.repo.CountUsers(ctx)
	if err != nil {
		return models.PaginatedUsersResponse{}, err
	}

	var response []models.UserWithAgeResponse
	for _, user := range users {
		response = append(response, models.UserWithAgeResponse{
			ID:   user.ID,
			Name: user.Name,
			Dob:  user.Dob.Format("2006-01-02"),
			Age:  CalculateAge(user.Dob),
		})
	}

	return models.PaginatedUsersResponse{
		Users: response,
		Page:  page,
		Limit: limit,
		Total: total,
	}, nil
}