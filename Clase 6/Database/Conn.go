package Database

import (
	"Ejemplo1/Instance"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func Connect() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error al cargar variable de entorno")
	}

	server := os.Getenv("DB_HOST")
	const dbName = "DB"
	var mongoURI = "mongodb://" + server + ":27017/" + dbName

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	Instance.Mg = Instance.MongoInstance(MongoInstance{
		Client: client,
		Db:     db,
	})

	return nil
}
