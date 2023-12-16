package entity

import (
	"ddd-demo/circle/notify"
	"ddd-demo/circle/value"

	"github.com/google/uuid"
)

// サークル　これはライフサイクルを持つのでエンティティ
type Circle struct {
	id      *value.CircleId
	name    *value.CircleName
	ownerId uuid.UUID
	members []uuid.UUID
}

func NewCircle(id *value.CircleId, name *value.CircleName, ownerId uuid.UUID, members []uuid.UUID) *Circle {
	return &Circle{
		id,
		name,
		ownerId,
		members,
	}
}

// 識別子はゲッターを公開
func (c *Circle) Id() *value.CircleId {
	return c.id
}

// それ以外は通知オブジェクトを使う
func (c *Circle) Notify() *notify.CircleNotification {
	return &notify.CircleNotification{
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
