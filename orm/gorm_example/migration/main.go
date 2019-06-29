package main

import (
	"github.com/JPMike/Starter-Golang/orm/gorm_example/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/gormigrate.v1"
	"log"
)

func main() {
	//db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("postgres", "postgres://gorm:password@localhost/gorm?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)
	// no table(s)
	db.SingularTable(true)

	// reference: https://gorm.io/docs/migration.html
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "201906281430",
			Migrate: func(db *gorm.DB) error {
				type Person struct {
					gorm.Model
					Name string
				}
				return db.AutoMigrate(&Person{}).Error
			},
			Rollback: func(db *gorm.DB) error {
				return db.DropTable("person").Error
			},
		},
		{
			ID: "201906281445",
			Migrate: func(db *gorm.DB) error {
				type Person struct {
					Age int `gorm:"index;not null"`
				}
				return db.AutoMigrate(&Person{}).Error
			},
			Rollback: func(db *gorm.DB) error {
				return db.Table("person").DropColumn("age").Error
			},
		},
		{
			ID: "201906281450",
			Migrate: func(db *gorm.DB) error {
				type Person struct {
					gorm.Model
					Name string
					Age  int `gorm:"index;not null"`
				}
				type Pet struct {
					gorm.Model
					Name     string
					Person   Person
					PersonID uint
				}
				if err := db.AutoMigrate(&Pet{}).Error; err != nil {
					return err
				}
				if err = db.Model(&model.Pet{}).AddForeignKey("person_id", "person(id)", "RESTRICT", "RESTRICT").Error; err != nil {
					// sqlite does not support ADD CONSTRAINT in ALTER TABLE
					return err
				}
				return nil
			},
			Rollback: func(db *gorm.DB) error {
				//if err := db.Model(&model.Pet{}).RemoveForeignKey("person_id", "person(id)").Error; err != nil {
				//	return err
				//}
				if err = db.DropTable("pet").Error; err != nil {
					return err
				}
				return nil
			},
		},
	})

	// if a clean db, will create the following tables, and then insert all migration ID in to migrations table
	m.InitSchema(func(db *gorm.DB) error {
		err := db.AutoMigrate(
			// all tables here
			&model.Person{},
			&model.Pet{},
		).Error
		if err != nil {
			return err
		}

		// all foreign keys here
		if err = db.Model(&model.Pet{}).AddForeignKey("person_id", "person(id)", "RESTRICT", "RESTRICT").Error; err != nil {
			// sqlite does not support ADD CONSTRAINT in ALTER TABLE
			return err
		}

		return nil
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")

	//if err = m.RollbackTo("201906281445"); err != nil {
	//	log.Fatalf("Could not migrate: %v", err)
	//}
	//log.Printf("RollbackTo did run successfully")

}
