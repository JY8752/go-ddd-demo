package entity

import (
	"ddd-demo/circle/value"
	"ddd-demo/user"

	"github.com/google/uuid"
)

// サークル　これはライフサイクルを持つのでエンティティ
type Circle struct {
	Id      *value.CircleId
	Name    *value.CircleName
	OwnerId uuid.UUID
	Members []user.User
}

func NewCircle(id *value.CircleId, name *value.CircleName, ownerId uuid.UUID, members []user.User) *Circle {
	return &Circle{
		id,
		name,
		ownerId,
		members,
	}
}
