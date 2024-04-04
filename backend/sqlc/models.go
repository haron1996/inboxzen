// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Deliverytime struct {
	DeliveryTime pgtype.Time        `json:"delivery_time"`
	DeliveryAmPm pgtype.Text        `json:"delivery_am_pm"`
	DateAdded    pgtype.Timestamptz `json:"date_added"`
	DateUpdated  pgtype.Timestamptz `json:"date_updated"`
	EmailAddress string             `json:"email_address"`
}

type Email struct {
	EmailAddress   string             `json:"email_address"`
	AccountName    string             `json:"account_name"`
	ProfilePicture string             `json:"profile_picture"`
	DateAdded      pgtype.Timestamptz `json:"date_added"`
	UserID         string             `json:"user_id"`
	Primaryaccount bool               `json:"primaryaccount"`
}

type User struct {
	ID           string             `json:"id"`
	RegisterDate pgtype.Timestamptz `json:"register_date"`
	LastLogin    pgtype.Timestamptz `json:"last_login"`
}

type Vipdomain struct {
	DomainName   string             `json:"domain_name"`
	DateAdded    pgtype.Timestamptz `json:"date_added"`
	EmailAddress string             `json:"email_address"`
}

type Vipemailaddress struct {
	VipEmailAddress string             `json:"vip_email_address"`
	DateAdded       pgtype.Timestamptz `json:"date_added"`
	EmailAddress    string             `json:"email_address"`
}

type Vipkeyword struct {
	Keyword      string             `json:"keyword"`
	DateAdded    pgtype.Timestamptz `json:"date_added"`
	EmailAddress string             `json:"email_address"`
}
