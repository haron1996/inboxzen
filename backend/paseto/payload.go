package paseto

import (
	"time"
)

type PayLoad struct {
	UserID   string    `json:"account_id"`
	Email    string    `json:"email"`
	IssuedAt time.Time `json:"issued_at"`
	Expiry   time.Time `json:"expiry"`
}

func newPayload(userID, email string, issuedAt, expiry time.Time) *PayLoad {
	return &PayLoad{
		UserID:   userID,
		Email:    email,
		IssuedAt: issuedAt,
		Expiry:   expiry,
	}
}
