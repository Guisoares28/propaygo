package main

import (
	"fmt"
	"propay/internal/config"
	"propay/internal/controller"
	"propay/internal/middleware"
	"propay/internal/repository"
	"propay/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Arquivo .env não encontrado")
		return
	}

	config.ConnectDatabase() //Chamando essa função as tabelas já serão criadas automaticamente pelo gorm.

	r := gin.Default()

	// Repositories
	userRepo := repository.NewUserRepository(config.DB)
	accountRepo := repository.NewAccountRepository(config.DB)

	//Services
	userService := service.NewUserService(userRepo)
	accountService := service.NewAccountService(accountRepo)
	tokenService, err := service.NewTokenService(userRepo)
	authService := service.NewAuthService(userRepo, tokenService)

	if err != nil {
		fmt.Println("Não foi possível iniciar o tokenService")
		return
	}

	//Controllers
	userController := controller.NewUserController(userService)
	accountController := controller.NewAccountController(accountService)
	authController := controller.NewAuthController(authService)

	//Rotas User
	r.POST("/user", userController.RegisterUser)

	//Rotas Account

	//Rotas Login
	r.POST("/login", authController.Login)

	//Grupo de rotas protegidas
	protected := r.Group("/auth")
	protected.Use(middleware.AuthMiddleware(tokenService))
	{
		protected.POST("/account", accountController.RegisterAccount)
		protected.POST("/account/:id", accountController.ChangeStatus)
	}

	r.Run(":8000")
}
