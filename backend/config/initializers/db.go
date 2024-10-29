package initializers

import (
	"context"
	"fmt"
	"log"

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
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Taipei",
		"localhost",
		"5432",
		"admin",
		"admin123",
		"mydb",
	)

	// 連接 PostgreSQL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// 測試連接
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("failed to get database instance: ", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("failed to ping database: ", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")
}

func ConnectToRedis() {
	ctx := context.Background()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
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
