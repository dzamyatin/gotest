package command

import (
	"gorm.io/gorm"
)

type GormSchemaMigrationCommand struct {
	db       *gorm.DB
	entities []interface{}
}

func (g GormSchemaMigrationCommand) GetCode() string {
	return "gorm_schema_migration"
}

func (g GormSchemaMigrationCommand) Execute(input Input) Output {
	err := g.db.AutoMigrate(g.entities...)

	if err != nil {
		return Output{
			code:    OutputCodeFail,
			message: err.Error(),
		}
	}

	return Output{
		code:    OutputCodeSuccess,
		message: "Done",
	}
}

func NewGormSchemaMigrationCommand(
	db *gorm.DB,
	entities []interface{},
) GormSchemaMigrationCommand {
	return GormSchemaMigrationCommand{
		db:       db,
		entities: entities,
	}
}
