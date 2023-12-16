package notify

import (
	"ddd-demo/circle/value"

	"github.com/google/uuid"
)

type CircleNotification struct {
	Name    *value.CircleName
	OwnerId uuid.UUID
	Members []uuid.UUID
}
