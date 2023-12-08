package users

import (
	"github.com/gofiber/fiber/v2"
)

func AddUserRoutes(fiberApp *fiber.App, userController *UserController) {
	userGroup := fiberApp.Group("/user")
	userGroup.Post("/create", userController.CreateUserHandler)
}

func Init(fiberApp *fiber.App, userStorage *UserStorage) {
	err := InitUsersTable(userStorage.userDB)
	if err != nil {
		panic(err)
	}
	userController := NewUserController(userStorage)
	AddUserRoutes(fiberApp, userController)
}
