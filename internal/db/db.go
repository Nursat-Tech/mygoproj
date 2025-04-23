package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/nursat/myproj/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbHost := "localhost"
	dbName := "postgres"
	dbUser := "postgres"
	dbPass := "51177477"
	dbPort := "5433"
	sslmode := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, sslmode)

	sqlDB, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, sslmode,
	))
	if err != nil {
		log.Fatal("SQL Open error: ", err)
	}

	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal("Migration driver error: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://C:/Users/Nursat/myproj/internal/db/migrations",
		"postgres", driver,
	)
	if err != nil {
		log.Fatal("Migration instance error: ", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("Migration up error: ", err)
	}

	// GORM DB
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("GORM connect error: ", err)
	}

	// ⚠ GORM auto migration (жоғарғы сенімділік үшін)
	if err := gormDB.AutoMigrate(&models.Player{}); err != nil {
		log.Fatal("AutoMigrate error: ", err)
	}

	DB = gormDB
}
