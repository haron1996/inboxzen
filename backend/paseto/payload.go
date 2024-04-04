package paseto

import (
	"time"
)

type PayLoad struct {
	UserID   string    `json:"account_id"`
	IssuedAt time.Time `json:"issued_at"`
	Expiry   time.Time `json:"expiry"`
}

func newPayload(userID string, issuedAt, expiry time.Time) *PayLoad {
	return &PayLoad{
		UserID:   userID,
		IssuedAt: issuedAt,
		Expiry:   expiry,
	}
}
