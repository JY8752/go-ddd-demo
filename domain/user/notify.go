package user

import "github.com/google/uuid"

type Notification struct {
	Id   uuid.UUID
	Name *FullName
}
