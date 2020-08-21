package controllers

import (
	"context"
	"log"

	"github.com/gofiber/fiber"
	"github.com/vbanurag/go-fiber/helper"
	"github.com/vbanurag/go-fiber/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllUsers(c *fiber.Ctx) {
	var users []models.User

	//Connection mongoDB with helper class
	collection := helper.ConnectDB("users")

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, c)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var user models.User
		// & character returns the memory address of the following variable.
		err := cur.Decode(&user) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(users)
}

func AddUser(c *fiber.Ctx) {

	p := new(models.User)
	// we decode our body request params
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	// connect db
	collection := helper.ConnectDB("users")

	// insert our user model.
	result, err := collection.InsertOne(context.TODO(), p)

	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(result)
}
