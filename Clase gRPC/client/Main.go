package main

import (
	"context"
	pb "ejemplo_despliegue/grpcClient"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ctx = context.Background()

type Data struct {
	Album  string
	Year   string
	Artist string
	Ranked string
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

	go sendRedisServer(rank)
	go sendMysqlServer(rank)

	return nil
}

func sendRedisServer(rank Data) {
	conn, err := grpc.Dial("", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Artist: rank.Artist,
		Album: rank.Album,
		Year: rank.Year,
		Ranked: rank.Ranked,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Respuesta del server " + ret.GetInfo())
}

func sendMysqlServer(rank Data) {

}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"res": "todo bien",
		})
	})
	app.Post("/insert", insertData)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
