package models

import "gorm.io/gorm"

type User struct {
	ID       int
	Name     string
	Age      int
	IsVerify bool   `gorm:"column:trusted"`
	Cards    []Card `gorm:"foreignKey:UserID"`
}

type Card struct {
	ID     int
	Number string
	Type   string
	UserID int
}

func InitModels(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Card{})
	if err != nil {
		return err
	}

	return nil
}
