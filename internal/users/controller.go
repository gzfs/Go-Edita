package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	userStorage *UserStorage
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserController(userStorage *UserStorage) *UserController {
	return &UserController{
		userStorage,
	}
}

func (userController *UserController) CreateUserHandler(fiberCtx *fiber.Ctx) error {
	createUserRequest := new(CreateUserRequest)
	if err := fiberCtx.BodyParser(createUserRequest); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to Create User",
		})
	}
	err := userController.userStorage.Create(&User{
		ID:    uuid.New().String(),
		Name:  createUserRequest.Name,
		Email: createUserRequest.Email,
	})

	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return fiberCtx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully Created User",
	})
}
