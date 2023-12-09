package users

import (
	"gzfs/Go~Edita/config"

	"github.com/gofiber/fiber/v2"
)

func AddUserRoutes(fiberApp *fiber.App, userController *UserController) {
	userGroup := fiberApp.Group("/user")
	userGroup.Post("/create", userController.Create)
	userGroup.Post("/", userController.GetByEmail)
	userGroup.Post("/login", userController.Login)
}

func Init(fiberApp *fiber.App, userStorage *UserStorage, envConfig config.Env) {
	err := InitUsersTable(userStorage.userDB)
	if err != nil {
		panic(err)
	}
	userController := NewUserController(userStorage, envConfig)
	AddUserRoutes(fiberApp, userController)
}
