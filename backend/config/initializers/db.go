package initializers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
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
		fmt.Printf("Database '%s' created successfully!\n", dbName)
	} else {
		fmt.Printf("Database '%s' already exists.\n", dbName)
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

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	fmt.Printf("Connected to '%s' database successfully!\n", dbName)
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

	val, err := Rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Fatal("Could not get key:", err)
	}
	fmt.Println("key:", val)
	fmt.Println("Connected to Redis successfully!")
}
