package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/go-redis/redis/v8"
)

var (
	PostgresDB *gorm.DB
	RedisClient *redis.Client
)

func ConnectPostgres(dsn string) {
	var err error
	PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to PostgreSQL: " + err.Error())
	}
}

func ConnectRedis(addr, password string) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
}

func Close() {
	sqlDB, _ := PostgresDB.DB()
	sqlDB.Close()
	RedisClient.Close()
}