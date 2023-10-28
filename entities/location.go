package entities

import "time"

type UserLocation struct {
	Id         int64   `json:"id,omitempty"`
	Longtitude float64 `json:"longitude,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (UserLocation) TableName() string {
	return "user_location"
}
