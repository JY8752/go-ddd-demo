package circle

import (
	"context"
	"database/sql"
	"ddd-demo/domain/circle"
	"ddd-demo/infrastructure"

	"github.com/google/uuid"
)

type circleRepository struct {
	db      *sql.DB
	queries *infrastructure.Queries
}

func NewCircleRepository(db *sql.DB) circle.CircleRepository {
	return &circleRepository{db, infrastructure.New(db)}
}

func (cr *circleRepository) Save(ctx context.Context, circle *circle.Circle) error {
	notification := circle.Notify()
	params := infrastructure.CreateCircleParams{
		ID:     circle.Id().Value(),
		Name:   notification.Name.Value(),
		UserID: notification.OwnerId.String(),
	}
	_, err := cr.queries.CreateCircle(ctx, params)
	return err
}

func (cr *circleRepository) FindByName(ctx context.Context, name *circle.CircleName) (*circle.Circle, error) {
	result, err := cr.queries.FindByCircleName(ctx, name.Value())
	if err != nil {
		return nil, err
	}
	return toEntity(result)
}

func toEntity(c infrastructure.Circle) (*circle.Circle, error) {
	id := circle.NewCircleId(c.ID)
	name, err := circle.NewCircleName(c.Name)
	if err != nil {
		return nil, err
	}
	ownerId := uuid.MustParse(c.UserID)
	return circle.NewCircle(id, name, ownerId, []uuid.UUID{}), nil
}
