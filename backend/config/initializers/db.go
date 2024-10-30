package initializers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Rdb *redis.Client
)

func ConnectToDB() {
	var err error
	var dbName string
	var dbHost string

	dbPort := "5432"
	dbUser := "admin"
	dbPassword := "admin123"

	if os.Getenv("ENV") == "docker" {
		dbHost = "postgres"
		dbName = "production"
	} else {
		dbHost = "localhost"
		dbName = "development"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable TimeZone=Asia/Taipei",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
	}

	var exists bool
	err = DB.Raw("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = ?)", dbName).Scan(&exists).Error
	if err != nil {
		log.Fatalf("failed to check if database exists: %v", err)
	}

	if !exists {
		if err := DB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error; err != nil {
			log.Fatalf("failed to create database: %v", err)
		}
		log.Printf("Database '%s' created successfully!\n", dbName)
	} else {
		log.Printf("Database '%s' already exists.\n", dbName)
	}

	dsn = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Taipei",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to %s database: %v", dbName, err)
	}

	log.Printf("Connected to '%s' database successfully!\n", dbName)

	applyMigrations()

}

func applyMigrations() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	driver, err := migratePostgres.WithInstance(sqlDB, &migratePostgres.Config{})
	if err != nil {
		log.Fatalf("could not create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../backend/db/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("migration setup failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("Migrations applied successfully!")
}

func ConnectToRedis() {
	var redisURL string
	var rDbNumber int

	ctx := context.Background()
	if os.Getenv("ENV") == "docker" {
		redisURL = "redis:6379"
		rDbNumber = 1
	} else {
		redisURL = "localhost:6379"
		rDbNumber = 0
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       rDbNumber,
	})

	err := Rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatal("Could not connect to Redis:", err)
	}

	err = Rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Fatal("Could not set key:", err)
	}

	log.Println("Connected to Redis successfully!")
}
