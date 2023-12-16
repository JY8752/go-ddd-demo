package user

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	id   uuid.UUID
	name *FullName
}

func NewFromName(name *FullName) *User {
	uuid := uuid.New()
	return &User{id: uuid, name: name}
}

func New(id uuid.UUID, name *FullName) *User {
	return &User{id, name}
}

// ドメインルールの漏洩につながるので不用意にゲッターを作るべきでない。
// func (u *User) Id() uuid.UUID {
// 	return u.id
// }

// func (u *User) Name() *FullName {
// 	return u.name
// }

// ゲッターのかわりに通知オブジェクトを使う
func (u *User) Notify() *UserNotification {
	return &UserNotification{u.id, u.name}
}

func (u *User) Equals(other *User) bool {
	return u.id == other.id
}

func (u *User) ChangeName(name *FullName) {
	u.name = name
}

func (u *User) String() string {
	return fmt.Sprintf("id: %s, name: %s", u.id.String(), u.name.String())
}
