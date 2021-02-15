package main

import (
	"context"
	"encoding/json"

	fiber "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	_id       string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int    `json:"age,omitempty"`
}

func GetPerson(c *fiber.Ctx) error {
	collection, err := GetMongoDBCollection(DBName, collectionName)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	defer cur.Close(context.Background())

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(fiber.StatusNotFound)
		return nil
	}

	json, _ := json.Marshal(results)
	c.Send(json)

	return nil
}

func CreatePerson(c *fiber.Ctx) error {
	collection, err := GetMongoDBCollection(DBName, collectionName)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}

	var person Person
	json.Unmarshal([]byte(c.Body()), &person)

	res, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}

	response, _ := json.Marshal(res)
	c.Send(response)

	return nil
}

func UpdatePerson(c *fiber.Ctx) error {
	collection, err := GetMongoDBCollection(DBName, collectionName)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	var person Person
	json.Unmarshal([]byte(c.Body()), &person)

	update := bson.M{
		"$set": person,
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}

	response, _ := json.Marshal(res)
	c.Send(response)

	return nil
}

func DeletePerson(c *fiber.Ctx) error {
	collection, err := GetMongoDBCollection(DBName, collectionName)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}

	jsonResponse, _ := json.Marshal(res)
	c.Send(jsonResponse)

	return nil
}
