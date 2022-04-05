package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbOrm struct {
	Db *gorm.DB
	M  *MigrateOrm
}

func NewDbOrm(dsn string, l logger.LogLevel) (*DbOrm, error) {
	orm := &DbOrm{}

	err := orm.connect(dsn, l)
	if err != nil {
		return nil, err
	}

	return orm, nil
}

func (o *DbOrm) connect(dsn string, l logger.LogLevel) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(l),
	})
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	o.Db = db
	o.M = NewMigrateOrm(o.Db)

	return nil
}

func (o *DbOrm) Migrate() error {
	err := o.M.Migrate()
	return err
}
