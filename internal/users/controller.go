package users

import (
	"fmt"
	"gzfs/Go~Edita/config"
	"gzfs/Go~Edita/pkg/helper"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	userStorage *UserStorage
	envConfig   config.Env
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Username   string `json:"username"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type GetUserRequest struct {
	Email string `json:"email"`
}

func NewUserController(userStorage *UserStorage, envConfig config.Env) *UserController {
	return &UserController{
		envConfig:   envConfig,
		userStorage: userStorage,
	}
}

func (userController *UserController) Login(fiberCtx *fiber.Ctx) error {
	loginRequest := new(LoginRequest)
	if err := fiberCtx.BodyParser(loginRequest); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to Login",
		})
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to Login",
		})
	}

	user, err := userController.userStorage.GetByEmail(loginRequest.Email)

	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to Login",
		})
	}

	if user.Email == "" {
		return fiberCtx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Failed to Login",
		})
	}

	_, err = helper.VerifyPassword(user.Password, loginRequest.Password)

	if err != nil {
		return fiberCtx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Failed to Login",
		})
	}

	signedToken, signedRefreshToken, err := helper.GenerateAllTokens(userController.envConfig, user.Email, user.First_Name, user.Last_Name, user.Admin)

	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to Login",
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":              "Successfully Logged In",
		"token":                signedToken,
		"refresh_token":        signedRefreshToken,
		"token_expiry":         time.Now().Add(time.Hour * 24).Format(time.RFC3339),
		"refresh_token_expiry": time.Now().Add(time.Hour * 24 * 7).Format(time.RFC3339),
	})
}

func (userController *UserController) GetByEmail(fiberCtx *fiber.Ctx) error {
	getUserRequest := new(GetUserRequest)
	if err := fiberCtx.BodyParser(getUserRequest); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to Get User",
		})
	}
	getUser, err := userController.userStorage.GetByEmail(getUserRequest.Email)
	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User not Found",
		})
	}
	return fiberCtx.Status(fiber.StatusOK).JSON(getUser)
}

func (userController *UserController) Create(fiberCtx *fiber.Ctx) error {
	createUserRequest := new(CreateUserRequest)
	if err := fiberCtx.BodyParser(createUserRequest); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to Create User",
		})
	}

	if createUserRequest.Username == "" || createUserRequest.First_Name == "" || createUserRequest.Last_Name == "" || createUserRequest.Email == "" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to Create User",
		})
	}

	hashedPassword, err := helper.HashPassword(createUserRequest.Password)

	if err != nil {
		fmt.Println(err)
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to Create User",
		})
	}

	err = userController.userStorage.Create(&User{
		ID:         uuid.New().String(),
		Email:      createUserRequest.Email,
		First_Name: createUserRequest.First_Name,
		Last_Name:  createUserRequest.Last_Name,
		Username:   createUserRequest.Username,
		Password:   hashedPassword,
		Created_At: time.Now(),
		Updated_At: time.Now(),
		Admin:      false,
	})

	if err != nil {
		fmt.Println(err)
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to Create User",
		})
	}

	return fiberCtx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully Created User",
	})
}
