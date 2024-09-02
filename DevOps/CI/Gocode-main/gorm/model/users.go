package model

type InfoUser struct {
	ID       int
	User     User     `json:"user" gorm:"foreignKey:UserId"`
	Messages Messages `json:"messages" gorm:"foreignKey:MessageID"`
}
type User struct {
	UserId  uint `json:"user_id"`
	Name    string
	Address string
	Tel     int
}

type Messages struct {
	MessageID uint `json:"message_id"`
	Age       int
	Stature   int64
}
