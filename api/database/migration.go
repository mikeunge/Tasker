package database

import (
	"github.com/mikeunge/Tasker/api/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Task{})
}

func CreateTable(db *gorm.DB, entity struct{}) error {
	return db.Migrator().CreateTable(&entity)
}

func TableExists(db *gorm.DB, entity struct{}) bool {
	return db.Migrator().HasTable(&entity)
}
