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

type SignupRequest struct {
	Email    string
	Password string
	Phone    string
}

func CreateUser(c *fiber.Ctx) error {
	var DB = database.ConnectDatabase()
	// usersCollection := collections.GetUsersCollections(DB, os.Getenv("UsersCollectionName"))
	usersCollection := collections.GetUsersCollections(DB, "Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(model.User)
	defer cancel()

	req := new(SignupRequest)
	err := c.BodyParser(req)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	log.Println(req)

	if req.Email == "" || req.Password == "" || req.Phone == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Signup Data")
	}

	salltedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Fatal(err)
	}

	user = &model.User{
		Email:     req.Email,
		Password:  salltedPassword,
		Phone:     req.Phone,
		CreatedAt: time.Now()}

	log.Println(user)
	result, err := usersCollection.InsertOne(ctx, user)
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
