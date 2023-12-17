package circle

import (
	"context"
	"ddd-demo/domain"
	"errors"

	"github.com/google/uuid"
)

type CircleRepository interface {
	Save(ctx context.Context, circle *Circle) error
	FindByName(ctx context.Context, name *CircleName) (*Circle, error)
}

// サークルの識別子　値オブジェクト
type CircleId struct {
	domain.ValueObject[string]
}

func NewCircleId(v string) *CircleId {
	return &CircleId{domain.NewValueObject[string](v)}
}

// サークル名 値オブジェクト
type CircleName struct {
	domain.ValueObject[string]
}

func NewCircleName(v string) (*CircleName, error) {
	nameLength := len([]rune(v))

	if nameLength < 3 {
		return nil, errors.New("circle name must be at least 3 charcters")
	}

	if nameLength > 20 {
		return nil, errors.New("circle name must be less than 20 charcters")
	}

	return &CircleName{domain.NewValueObject[string](v)}, nil
}

// サークル　これはライフサイクルを持つのでエンティティ
type Circle struct {
	id      *CircleId
	name    *CircleName
	ownerId uuid.UUID
	members []uuid.UUID
}

func NewCircle(id *CircleId, name *CircleName, ownerId uuid.UUID, members []uuid.UUID) *Circle {
	return &Circle{
		id,
		name,
		ownerId,
		members,
	}
}

// 識別子はゲッターを公開
func (c *Circle) Id() *CircleId {
	return c.id
}

// それ以外は通知オブジェクトを使う
func (c *Circle) Notify() *CircleNotification {
	return &CircleNotification{
		Name:    c.name,
		OwnerId: c.ownerId,
		Members: c.members,
	}
}

// const maxMembersCount = 30
// func (c *Circle) IsFull() bool {
// 	// return len(c.members) > 29
// 	return len(c.members) + 1 > maxMembersCount
// }

func (c *Circle) CountMembers() int {
	return len(c.members) + 1
}
