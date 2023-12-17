package circle

import (
	"ddd-demo/domain/user"

	"github.com/google/uuid"
)

func Create(name *CircleName, owner *user.User) *Circle {
	id := uuid.NewString()
	return NewCircle(
		NewCircleId(id),
		name,
		owner.Notify().Id,
		[]uuid.UUID{},
	)
}
