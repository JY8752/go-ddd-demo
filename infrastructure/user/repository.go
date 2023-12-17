package user

import (
	"context"
	"database/sql"
	"ddd-demo/domain/user"
	"ddd-demo/infrastructure"

	"github.com/google/uuid"
)

type repository struct {
	db      *sql.DB
	queries *infrastructure.Queries
}

func NewRepository(db *sql.DB) user.Repository {
	return &repository{db, infrastructure.New(db)}
}

func (r *repository) Create(ctx context.Context, user *user.User) (*user.User, error) {
	notification := user.Notify()
	params := infrastructure.CrateUserParams{
		ID:        notification.Id.String(),
		FirstName: notification.Name.FirstName(),
		LastName:  notification.Name.LastName(),
	}

	u, err := r.queries.CrateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return convertEntity(u)
}

func (r *repository) FindById(ctx context.Context, id string) (*user.User, error) {
	u, err := r.queries.FindUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertEntity(u)
}

func (r *repository) FindByName(ctx context.Context, firstName, lastName string) (*user.User, error) {
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

func (r *repository) UpdateName(ctx context.Context, user *user.User) (*user.User, error) {
	notification := user.Notify()
	params := infrastructure.UpdateUserNameParams{
		FirstName: notification.Name.FirstName(),
		LastName:  notification.Name.LastName(),
		ID:        notification.Id.String(),
	}

	u, err := r.queries.UpdateUserName(ctx, params)
	if err != nil {
		return nil, err
	}

	return convertEntity(u)
}

func (r *repository) Delete(ctx context.Context, id string) (*user.User, error) {
	u, err := r.queries.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertEntity(u)
}

func convertEntity(u infrastructure.User) (*user.User, error) {
	id := uuid.MustParse(u.ID)
	name, err := user.NewFullName(u.FirstName, u.LastName)
	if err != nil {
		return nil, err
	}

	return user.New(id, name), nil
}
