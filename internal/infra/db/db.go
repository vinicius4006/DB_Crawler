package infra

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type db struct {
	DB *gorm.DB
}

func NewDB() (db, error) {
	databaseName := "postgres"
	userName := "postgres"
	password := "ifma2023"
	host := "localhost"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable timeZone=America/Sao_Paulo",
		host, userName, password, databaseName)

	fmt.Println(dsn)

	//Open connection to a postgresql database:
	dbOrm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	return db{DB: dbOrm}, err
}
