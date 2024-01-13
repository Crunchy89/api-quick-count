package responses

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	UUID      uuid.UUID `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
