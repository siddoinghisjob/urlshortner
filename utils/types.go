package utils

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type message struct {
	Text  string `json:"message"`
	Error bool   `json:"error"`
}

type urlLink struct {
	URL string `json:"URL"`
}

type notFoundError struct {
	message string
}

type internalError struct {
	message string
}

func dotENV(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return os.Getenv(key)
}

func dotENVInt(key string) int {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	i, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(err)
	}
	return i
}

var ctx = context.Background()

type analyticsStore struct {
	url        string
	name, date string
}

type analyticsLogger struct {
	mutex  sync.Mutex
	db     *sql.DB
	buffer []analyticsStore
}

type analyticsData struct {
	Name string `json:"name"`
	Data int    `json:"data"`
}

var Client = redis.NewClient(&redis.Options{
	Addr:     dotENV("REDIS_ADDR"),
	Password: dotENV("REDIS_PASSWORD"),
	DB:       dotENVInt("REDIS_DB"),
})

var (
	host     = dotENV("PG_HOST")
	port     = dotENVInt("PG_PORT")
	user     = dotENV("PG_USER")
	password = dotENV("PG_PASSWORD")
	dbname   = dotENV("PG_DB")
)

var SQLInfo string = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

var DB, DBErr = sql.Open("postgres", SQLInfo)
