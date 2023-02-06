package auth

import (
	"context"
	"fiber/collections"
	"fiber/database"
	"fiber/model"
	"fiber/utils"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var DB = database.ConnectDatabase()
	// usersCollection := collections.GetUsersCollections(DB, os.Getenv("UsersCollectionName"))
	usersCollection := collections.GetUsersCollections(DB, "Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(model.User)
	defer cancel()

	err := c.BodyParser(user)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	salltedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}

	postPayload := model.User{
		Email:     user.Email,
		Password:  salltedPassword,
		Phone:     user.Phone,
		CreatedAt: time.Now()}

	log.Println(postPayload)
	result, err := usersCollection.InsertOne(ctx, postPayload)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}
	c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "Successfully Added User",
		"data":    result,
	})

	return nil
}

func login() {

}
