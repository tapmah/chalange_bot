package db

import (
	"fmt"
	"log"

	"github.com/tapmah/chalange_bot/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

var globalDb *Database

func getDb() *Database {
	if globalDb == nil {
		InitDB()
	}
	return globalDb
}

func InitDB() {
	cfg := config.GetConfig()
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	var err error
	globalDb, err = NewDatabase(source)
	if err != nil {
		log.Fatalf("error creat db")
	}
}

func NewDatabase(dataSourceName string) (*Database, error) {
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}

	return &Database{Conn: db}, nil
}

func (d *Database) AddUser(username string, telegramId int64) error {
	user := User{Username: username, TelegramId: telegramId}
	return d.Conn.Create(&user).Error
}

func (d *Database) UserExists(username string) (bool, error) {
	var count int64
	err := d.Conn.Model(&User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}
