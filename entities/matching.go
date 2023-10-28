package entities

import "time"

type UserMatching struct {
	UserId    int64 `json:"user_id,omitempty"`
	PartnerId int64 `json:"partner_id,omitempty"`
	IsLike    int   `json:"is_like,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (UserMatching) TableName() string {
	return "matching"
}
