package main

import (
	_ "github.com/DiegoJCordeiro/golang-study/chapter6/docs"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/handler"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/infra/database"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/infra/environment"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/infra/webserver/handlers"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/service"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"log"
)

// @title Postgraduate - GO Expert
// @version 1.0
// @description API for Go Expert Postgraduate.
// @termsOfService http://swagger.io/terms/

// @contact.name Diego Cordeiro
// @contact.url https://github.com/DiegoJCordeiro
// @contact.email diegocordeiro.contatos@gmail.com

// @license.name Diego Cordeiro License
// @license.url  https://github.com/DiegoJCordeiro/golang-study/blob/main/LICENSE

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	var userDB database.UserDB = database.NewUserDB(environment.GetDBPostgres())
	var userService service.UserService = service.NewUserService(userDB)
	var userHandler = handlers.NewUserHandler(userService)

	var tokenService = service.NewTokenService(userDB)
	var tokenHandler handlers.TokenHandler = handlers.NewTokenHandler(tokenService)

	var roleDB database.RoleDB = database.NewRoleDB(environment.GetDBPostgres())
	var roleService service.RoleService = service.NewRoleService(roleDB)
	var roleHandler handlers.RoleHandler = handlers.NewRoleHandler(roleService)

	fiberServer := fiber.New(fiber.Config{
		ErrorHandler: handler.NewErrorAPI().ErrorHandler,
	})

	userApi := fiberServer.Group("/user")
	userApi.Post("/", userHandler.CreateUser)
	userApi.Post("/token", tokenHandler.GenerateToken)

	fiberServer.Get("/docs/*", swagger.New(swagger.Config{
		URL: "/docs/doc.json",
	}))

	fiberServer.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(environment.GetJwtSecret())},
	}))

	userApi.Get("/findByName", userHandler.GetUserByName)

	roleApi := fiberServer.Group("/role")
	roleApi.Post("/", roleHandler.CreateRole)
	roleApi.Put("/", roleHandler.UpdateRole)
	roleApi.Delete("/", roleHandler.DeleteRole)
	roleApi.Get("/", roleHandler.GetAllRole)
	roleApi.Get("/findByName", roleHandler.GetRoleByName)

	errServerMux := fiberServer.Listen(":8080")

	if errServerMux != nil {
		log.Fatal(errServerMux)
	}
}
