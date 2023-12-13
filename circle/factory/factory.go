package factory

import (
	"ddd-demo/circle/entity"
	"ddd-demo/circle/value"
	"ddd-demo/user"

	"github.com/google/uuid"
)

func CreateCircle(name *value.CircleName, owner *user.User) *entity.Circle {
	id := uuid.NewString()
	return &entity.Circle{
		Id:      value.NewCircleId(id),
		Name:    name,
		OwnerId: owner.Id(),
	}
}
