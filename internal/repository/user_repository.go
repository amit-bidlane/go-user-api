package repository

import (
	"context"
	"database/sql"

	db "go-user-api/db/sqlc"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{
		queries: db.New(database),
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, params db.CreateUserParams) (db.User, error) {
	return r.queries.CreateUser(ctx, params)
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

func (r *UserRepository) UpdateUser(ctx context.Context, params db.UpdateUserParams) (db.User, error) {
	return r.queries.UpdateUser(ctx, params)
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}

func (r *UserRepository) ListUsers(ctx context.Context, params db.ListUsersParams) ([]db.User, error) {
	return r.queries.ListUsers(ctx, params)
}

func (r *UserRepository) CountUsers(ctx context.Context) (int64, error) {
	return r.queries.CountUsers(ctx)
}