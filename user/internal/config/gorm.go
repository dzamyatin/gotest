package config

import (
	"app/user/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetGorm() *gorm.DB {
	return syncGetOrCreateByType(
		func() *gorm.DB {
			conf := GetDatabaseConfig()
			db, err := gorm.Open(sqlite.Open(conf.Path), &gorm.Config{})
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
