package controllers

import (
	"context"
	"log"

	"github.com/gofiber/fiber"
	"github.com/vbanurag/go-fiber/helper"
	"github.com/vbanurag/go-fiber/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetUser(c *fiber.Ctx) {

	var user models.User

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(c.Params("id"))

	collection := helper.ConnectDB("users")

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	c.JSON(user)
}

func EditUser(c *fiber.Ctx) {

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(c.Params("id"))

	var user models.User

	p := new(models.User)

	collection := helper.ConnectDB("users")

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	// prepare update model.
	update := bson.M{
		"$set": bson.M{
			"name": bson.M{
				"firstname": p.Name.FirstName,
				"lastname":  p.Name.LastName,
			},
		},
	}
	upsert := true
	after := options.After

	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&user)

	if err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) {

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(c.Params("id"))

	collection := helper.ConnectDB("users")

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	c.JSON(deleteResult)
}
