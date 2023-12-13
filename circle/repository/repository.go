package repository

import (
	"context"
	"database/sql"
	"ddd-demo/circle/entity"
	"ddd-demo/circle/value"
	"ddd-demo/infrastructure"
	"ddd-demo/user"

	"github.com/google/uuid"
)

type CircleRepository interface {
	Save(ctx context.Context, circle *entity.Circle) error
	FindByName(ctx context.Context, name *value.CircleName) (*entity.Circle, error)
}

type circleRepository struct {
	db      *sql.DB
	queries *infrastructure.Queries
}

func NewCircleRepository(db *sql.DB) CircleRepository {
	return &circleRepository{db, infrastructure.New(db)}
}

func (cr *circleRepository) Save(ctx context.Context, circle *entity.Circle) error {
	params := infrastructure.CreateCircleParams{
		ID:     circle.Id.Value(),
		Name:   circle.Name.Value(),
		UserID: circle.OwnerId.String(),
	}
	_, err := cr.queries.CreateCircle(ctx, params)
	return err
}

func (cr *circleRepository) FindByName(ctx context.Context, name *value.CircleName) (*entity.Circle, error) {
	result, err := cr.queries.FindByCircleName(ctx, name.Value())
	if err != nil {
		return nil, err
	}
	return toEntity(result)
}

func toEntity(c infrastructure.Circle) (*entity.Circle, error) {
	id := value.NewCircleId(c.ID)
	name, err := value.NewCircleName(c.Name)
	if err != nil {
		return nil, err
	}
	ownerId := uuid.MustParse(c.UserID)
	return entity.NewCircle(id, name, ownerId, []user.User{}), nil
}
