package user

import (
	"context"
	"database/sql"
	"ddd-demo/infrastructure"
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

func (r *Repository) CreateUser(ctx context.Context, id, firstName, lastName string) (*DTO, error) {
	params := infrastructure.CrateUserParams{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	}

	u, err := r.queries.CrateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return convertDTO(u), nil
}

func (r *Repository) FindById(ctx context.Context, id string) (*DTO, error) {
	u, err := r.queries.FindUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertDTO(u), nil
}

func (r *Repository) FindByName(ctx context.Context, firstName, lastName string) (*DTO, error) {
	params := infrastructure.FindUserByNameParams{
		FirstName: firstName,
		LastName:  lastName,
	}

	u, err := r.queries.FindUserByName(ctx, params)
	if err != nil {
		return nil, err
	}

	return convertDTO(u), nil
}

func (r *Repository) UpdateName(ctx context.Context, id, firstName, lastName string) (*DTO, error) {
	params := infrastructure.UpdateUserNameParams{
		FirstName: firstName,
		LastName:  lastName,
		ID:        id,
	}

	u, err := r.queries.UpdateUserName(ctx, params)
	if err != nil {
		return nil, err
	}

	return convertDTO(u), nil
}

func (r *Repository) Delete(ctx context.Context, id string) (*DTO, error) {
	u, err := r.queries.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertDTO(u), nil
}

func convertDTO(u infrastructure.User) *DTO {
	return NewDTO(u.ID, u.FirstName, u.LastName)
}
