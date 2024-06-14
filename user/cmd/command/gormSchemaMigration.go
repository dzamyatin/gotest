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
	return Output{code: OutputCodeSuccess}
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
