package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

var Database *gorm.DB
var Uri = "root:root123@tcp(35.202.95.255)/test?charset=utf8mb4&parseTime=True&loc=Local"

var ctx = context.Background()
var rdb *redis.Client

type Data struct {
	gorm.Model
	Album  string
	Year   string
	Artist string
	Ranked string
}

func mysqlConnect() error {
	var err error
	Database, err = gorm.Open(mysql.Open(Uri), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = Database.AutoMigrate(&Data{})
	if err != nil {
		return err
	}

	return nil
}

func redisConnect() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       15,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pong)
}

func insertData(c *fiber.Ctx) error {
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	rank := Data{
		Album:  data["album"],
		Year:   data["year"],
		Artist: data["artist"],
		Ranked: data["ranked"],
	}

	go insertRedis(rank)
	go insertMysql(rank)

	return nil
}

func insertRedis(rank Data) {
	array := rank.Artist + "-" + rank.Year
	ranked, _ := strconv.ParseFloat(rank.Ranked, 64)

	rdb.ZAddArgs(ctx, array, redis.ZAddArgs{
		NX: true,
		Members: []redis.Z{{
			Score:  ranked,
			Member: rank.Album,
		}},
	})

	key := array + "-" + rank.Album
	rdb.HIncrBy(ctx, key, rank.Ranked, 1)

	fmt.Println(rank)
}

func insertMysql(rank Data) {
	Database.Create(&rank)
}

func main() {
	app := fiber.New()

	redisConnect()
	err := mysqlConnect()
	if err != nil {
		return
	}

	app.Post("/insert", insertData)

	err = app.Listen(":3000")
	if err != nil {
		return
	}
}
