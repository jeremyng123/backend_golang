package main

import (
	"context"
	"encoding/json"

	fiber "github.com/gofiber/fiber/v2"
)

func GetPerson(c *fiber.Ctx) {

}

func CreatePerson(c *fiber.Ctx) {
	collection, err := GetMongoDbCollection(dbName, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var person Person
	json.Unmarshal([]byte(c.Body()), &person)

	res, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}

func UpdatePerson(c *fiber.Ctx) {

}

func DeletePerson(c *fiber.Ctx) {

}
