package cache

import (
	"fmt"
	"github.com/alireza-fa/blog-go/src/constants"
	"github.com/go-redis/redis"
	"log"
	"os"
	"strconv"
	"time"
)

var redisClient *redis.Client

func InitRedis() error {
	if os.Getenv("DEBUG") == "" {
		err := os.Setenv(constants.RedisHost, "localhost")
		if err != nil {
			panic("error while setting REDIS_HOST in .env")
		}
	}

	db, err := strconv.Atoi(os.Getenv(constants.RedisDB))
	if err != nil {
		panic(err)
	}

	dialTimeout, err := strconv.Atoi(os.Getenv(constants.RedisDialTimeout))
	if err != nil {
		panic(err)
	}

	readTimeout, err := strconv.Atoi(os.Getenv(constants.RedisReadTimeout))
	if err != nil {
		panic(err)
	}

	writeTimeout, err := strconv.Atoi(os.Getenv(constants.RedisWriteTimeout))
	if err != nil {
		panic(err)
	}

	poolSize, err := strconv.Atoi(os.Getenv(constants.RedisPoolSize))
	if err != nil {
		panic(err)
	}

	poolTimeout, err := strconv.Atoi(os.Getenv(constants.RedisPoolTimeout))
	if err != nil {
		panic(err)
	}

	idleCheckFrequency, err := strconv.Atoi(os.Getenv(constants.RedisIdleCheckFrequency))
	if err != nil {
		panic(err)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%s", os.Getenv(constants.RedisHost), os.Getenv(constants.RedisPort)),
		Password:           os.Getenv(constants.RedisPassword),
		DB:                 db,
		DialTimeout:        time.Duration(dialTimeout) * time.Second,
		ReadTimeout:        time.Duration(readTimeout) * time.Second,
		WriteTimeout:       time.Duration(writeTimeout) * time.Second,
		PoolSize:           poolSize,
		PoolTimeout:        time.Duration(poolTimeout),
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: time.Duration(idleCheckFrequency) * time.Millisecond,
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		return err
	}

	log.Println("connect to redis")

	return nil
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}
