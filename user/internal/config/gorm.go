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
			db, err := gorm.Open(sqlite.Open(conf.Path), &gorm.Config{
				FullSaveAssociations: true, //For scenarios where a full update of the associated data is required (not just the foreign key references)
			})
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
