package user

import (
	"context"
	"database/sql"
	"ddd-demo/infrastructure"

	"github.com/google/uuid"
)

type Repository struct {
	db      *sql.DB
	queries *infrastructure.Queries
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db, infrastructure.New(db)}
}

func (r *Repository) CreateTables(ctx context.Context, ddl string) error {
	_, err := r.db.ExecContext(ctx, ddl)
	return err
}

func (r *Repository) CreateUser(ctx context.Context, id, firstName, lastName string) (*User, error) {
	params := infrastructure.CrateUserParams{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	}

	u, err := r.queries.CrateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return convertEntity(u)
}

func (r *Repository) FindById(ctx context.Context, id string) (*User, error) {
	u, err := r.queries.FindUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertEntity(u)
}

func (r *Repository) FindByName(ctx context.Context, firstName, lastName string) (*User, error) {
	params := infrastructure.FindUserByNameParams{
		FirstName: firstName,
		LastName:  lastName,
	}

	u, err := r.queries.FindUserByName(ctx, params)
	if err != nil {
		return nil, err
	}

	return convertEntity(u)
}

func (r *Repository) UpdateName(ctx context.Context, id, firstName, lastName string) (*User, error) {
	params := infrastructure.UpdateUserNameParams{
		FirstName: firstName,
		LastName:  lastName,
		ID:        id,
	}

	u, err := r.queries.UpdateUserName(ctx, params)
	if err != nil {
		return nil, err
	}

	return convertEntity(u)
}

func (r *Repository) Delete(ctx context.Context, id string) (*User, error) {
	u, err := r.queries.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertEntity(u)
}

func convertEntity(u infrastructure.User) (*User, error) {
	id := uuid.MustParse(u.ID)
	name, err := NewFullName(u.FirstName, u.LastName)
	if err != nil {
		return nil, err
	}

	return New(id, name), nil
}
