package Controller

import (
	"Ejemplo1/Instance"
	"Ejemplo1/Model"
	"context"
	"log"
)

func InsertData(nameColl string, dataParam string) {
	collection := Instance.Mg.Db.Collection(nameColl)
	doc := Model.Data{Percent: dataParam}

	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatalln(err)
	}
}
