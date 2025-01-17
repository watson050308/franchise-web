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
	DB     *gorm.DB
	Rdb    *redis.Client
	Getenv = os.Getenv
)

func GetCurrentEnv() string {
	env := Getenv("ENV")
	if env == "" {
		return "development"
	} else if env == "test" {
		return "test"
	}
	log.Println("ENV:", env)
	return env
}

func ConnectToDB() error {
	var err error
	var dbName, dbHost string

	dbPort := "5432"
	dbUser := "admin"
	dbPassword := "admin123"
	eenv := Getenv("ENV")
	log.Println("eenv: ", eenv)
	switch GetCurrentEnv() {
	case "test":
		dbHost = "localhost"
		dbName = "test_db"
	case "docker":
		dbHost = "postgres"
		dbName = "production"
	default:
		dbHost = "localhost"
		dbName = "development"
	}

	defaultDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable TimeZone=Asia/Taipei",

		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)

	log.Println("defaultDSN: ", defaultDSN)

	defaultDB, err := gorm.Open(postgres.Open(defaultDSN), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to default database: %w", err)
	}

	var exists bool
	err = defaultDB.Raw("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = ?)", dbName).Scan(&exists).Error
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}

	if !exists {
		log.Printf("Database '%s' does not exist. Creating...\n", dbName)
		if err := defaultDB.Exec(fmt.Sprintf("CREATE DATABASE \"%s\"", dbName)).Error; err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}
		log.Printf("Database '%s' created successfully!\n", dbName)
	} else {
		log.Printf("Database '%s' already exists.\n", dbName)
	}

	sqlDB, err := defaultDB.DB()
	if err == nil {
		sqlDB.Close()
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Taipei",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to %s database: %w", dbName, err)
	}

	log.Printf("Connected to '%s' database successfully!\n", dbName)

	err = ApplyMigrations(dbName)
	if err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}

func ApplyMigrations(dbName string) error {
	fileUrl := func() string {
		if dbName == "production" {
			return "file:///root/db/migrations"
		} else if dbName == "test_db" {
			return "file://../../../backend/db/migrations"
		}
		return "file://../backend/db/migrations"
	}()
	wd, _ := os.Getwd()
	log.Println(wd)
	log.Println("fileUrl: ", fileUrl)
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	driver, err := migratePostgres.WithInstance(sqlDB, &migratePostgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fileUrl,
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("migration setup failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %v", err)
	}

	log.Println("Migrations applied successfully!")
	return nil
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
