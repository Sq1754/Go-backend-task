package repository

import (
	"context"
	"time"

	db "github.com/sq1754/user-age-api/db/sqlc"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{
		q: q,
	}
}

// Create user
func (r *UserRepository) Create(
	ctx context.Context,
	name string,
	dob time.Time,
) (db.User, error) {
	return r.q.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

// Get user by ID
func (r *UserRepository) GetByID(
	ctx context.Context,
	id int32,
) (db.User, error) {
	return r.q.GetUserByID(ctx, id)
}

// List all users
func (r *UserRepository) List(
	ctx context.Context,
) ([]db.User, error) {
	return r.q.ListUsers(ctx)
}

// Update user
func (r *UserRepository) Update(
	ctx context.Context,
	id int32,
	name string,
	dob time.Time,
) (db.User, error) {
	return r.q.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

// Delete user
func (r *UserRepository) Delete(
	ctx context.Context,
	id int32,
) error {
	return r.q.DeleteUser(ctx, id)
}
