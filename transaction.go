package tavern

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	amount    float64
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
