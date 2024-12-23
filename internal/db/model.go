package db

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"unique;not null"`
	TelegramId int64  `gorm:"unique;not null"`
}
