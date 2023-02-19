package auth

import (
	"context"
	"fiber/collections"
	"fiber/database"
	"fiber/model"
	"fiber/utils"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupRequest struct {
	Email    string
	Password string
	Phone    string
}

type LoginRequest struct {
	Email    string
	Password string
}

func CreateUser(c *fiber.Ctx) error {
	/*
		Steps
		1. Conned
	*/
	var DB = database.ConnectDatabase()
	// usersCollection := collections.GetUsersCollections(DB, os.Getenv("UsersCollectionName"))
	usersCollection := collections.GetUsersCollections(DB, "Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userSignupRequestData := new(model.User)
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
	//log.Println(req)

	if req.Email == "" || req.Password == "" || req.Phone == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Signup Data")
	}

	salltedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Fatal(err)
	}

	userSignupRequestData = &model.User{
		Email:     req.Email,
		Password:  salltedPassword,
		Phone:     req.Phone,
		CreatedAt: time.Now()}

	// log.Println(userSignupRequestData)
	result, err := usersCollection.InsertOne(ctx, userSignupRequestData)
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

func Login(c *fiber.Ctx) error {
	req := new(LoginRequest)
	err := c.BodyParser(req)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	var DB = database.ConnectDatabase()
	usersCollection := collections.GetUsersCollections(DB, "Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var response primitive.M

	user, err := usersCollection.FindOne(ctx, bson.D{{"email", req.Email}}).Decode(&response)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
