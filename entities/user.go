package entities

import "time"

type Gender string

const (
	Male   Gender = "male"
	Female        = "female"
)

type User struct {
	Id          int64     `json:"id,omitempty" gorm:"primaryKey,autoIncrement"`
	Username    string    `json:"username,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Email       string    `json:"email,omitempty"`
	Age         uint      `json:"age,omitempty" gorm:"-"`
	BirthDay    time.Time `json:"birthday,omitempty" gorm:"column:birthday"`
	Gender      Gender    `json:"gender,omitempty"`
	Interesting string    `json:"interesting,omitempty"`
	IsActive    int       `json:"is_active,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u User) TableName() string {
	return "user"
}
