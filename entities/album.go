package entities

type Image struct {
	Id     int64  `json:"id,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
	Url    string `json:"url,omitempty"`
}

func (Image) TableName() string {
	return "album"
}
