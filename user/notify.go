package user

import "github.com/google/uuid"

type UserNotification struct {
	Id   uuid.UUID
	Name *FullName
}
