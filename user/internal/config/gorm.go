package config

import (
	"app/user/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetGorm() *gorm.DB {
	return syncGetOrCreateByType(
		func() *gorm.DB {
			db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
			if err != nil {
				panic("failed to connect database")
			}

			return db
		},
	)
}

func GetGormEntities() []interface{} {
	return []interface{}{
		entity.User{},
	}
}
