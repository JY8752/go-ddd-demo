package circle

import (
	"github.com/google/uuid"
)

type CircleNotification struct {
	Name    *CircleName
	OwnerId uuid.UUID
	Members []uuid.UUID
}
