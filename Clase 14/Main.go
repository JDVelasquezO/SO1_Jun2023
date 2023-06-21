package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

var ctx = context.Background()
var rdb *redis.Client

type Data struct {
	Album  string
	Year   string
	Artist string
	Ranked string
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

	fmt.Println(data)

	return nil
}

func main() {
	app := fiber.New()

	redisConnect()

	app.Post("/insert", insertData)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
