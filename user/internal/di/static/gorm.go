package static

import (
	"app/user/internal/di/singleton"
	"app/user/internal/entity"
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetGorm() *gorm.DB {
	return singleton.GlobalGetOrCreateTyped(
		func() *gorm.DB {
			//conf := config.GetConfig()

			//db, err := gorm.Open(sqlite.Open(conf.Path), &gorm.Config{
			//	ConnPool:             GetDB(), //Pool???
			//	FullSaveAssociations: true,    //For scenarios where a full update of the associated data is required (not just the foreign key references)
			//	PrepareStmt:          true,    //use prepared statmeent to speed up
			//	//SkipDefaultTransaction: true, //
			//})
			db, err := gorm.Open(
				sqlite.New(sqlite.Config{
					Conn: GetDB(),
				}),
				&gorm.Config{
					//ConnPool:             GetDB(), //Pool???
					FullSaveAssociations: true, //For scenarios where a full update of the associated data is required (not just the foreign key references)
					PrepareStmt:          true, //use prepared statmeent to speed up
					//SkipDefaultTransaction: true, //
				},
			)
			if err != nil {
				panic("failed to connect database")
			}

			return db
		},
	)
}

func NewGormSession(ctx context.Context) *gorm.DB {
	return GetGorm().Session(&gorm.Session{
		Context: ctx,
	})
}

func GetGormEntities() []interface{} {
	return []interface{}{
		entity.User{},
	}
}
