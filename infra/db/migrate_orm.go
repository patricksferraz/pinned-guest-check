package db

import (
	"fmt"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type MigrateOrm struct {
	Db *gorm.DB
	m  *gormigrate.Gormigrate
}

func NewMigrateOrm(db *gorm.DB) *MigrateOrm {
	m := MigrateOrm{
		Db: db,
	}
	m.load()
	return &m
}

func (m *MigrateOrm) load() {
	m.m = gormigrate.New(m.Db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202203301940",
			Migrate: func(db *gorm.DB) error {
				type Base struct {
					ID        *string    `gorm:"type:uuid;primaryKey"`
					CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
					UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
				}
				type Employee struct {
					Base
				}
				type Guest struct {
					Base
				}
				type Place struct {
					Base
				}
				type GuestCheckItem struct {
					Base
					Status         *int            `gorm:"column:status;not null"`
					CanceledReason *string         `gorm:"column:canceled_reason;type:varchar(255)"`
					Name           *string         `gorm:"column:name;not null"`
					Code           *int            `gorm:"column:code;not null"`
					Quantity       *int            `gorm:"column:quantity;not null"`
					UnitPrice      *float64        `gorm:"column:unit_price;not null"`
					Discount       *float64        `gorm:"column:discount"`
					TotalPrice     *float64        `gorm:"column:total_price"`
					FinalPrice     *float64        `gorm:"column:final_price;not null"`
					Note           *string         `gorm:"column:note;type:varchar(255)"`
					Tags           *pq.StringArray `gorm:"column:tags;type:text[]"`
					GuestCheckID   *string         `gorm:"column:guest_check_id;type:uuid;not null"`
				}
				type GuestCheck struct {
					Base
					TotalPrice     *float64 `gorm:"column:total_price"`
					TotalDiscount  *float64 `gorm:"column:total_discount"`
					FinalPrice     *float64 `gorm:"column:final_price"`
					Status         *int     `gorm:"column:status;not null"`
					CanceledReason *string  `gorm:"column:canceled_reason;type:varchar(255)"`
					Local          *string  `gorm:"column:local;type:varchar(255)"`
					Token          *string  `gorm:"column:token;type:varchar(25);not null"`
					GuestID        *string  `gorm:"column:guest_id;type:uuid;not null"`
					PlaceID        *string  `gorm:"column:place_id;type:uuid;not null"`
					AttendedBy     *string  `gorm:"column:attended_by;type:uuid"`
				}
				type Item struct {
					Base
					Code      *int            `gorm:"column:code;not null"`
					Name      *string         `gorm:"column:name;not null"`
					Available *bool           `gorm:"column:available;not null"`
					Price     *float64        `gorm:"column:price;not null"`
					Discount  *float64        `gorm:"column:discount"`
					Tags      *pq.StringArray `gorm:"column:tags;type:text[]"`
				}

				return db.AutoMigrate(&Employee{}, &Guest{}, &Place{}, &GuestCheckItem{}, &GuestCheck{}, &Item{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable("employees", "guests", "places", "guest_check_items", "guest_checks", "items")
			},
		},
	})
}

func (m *MigrateOrm) Migrate() error {
	if err := m.m.Migrate(); err != nil {
		return fmt.Errorf("could not migrate: %v", err)
	}
	return nil
}
