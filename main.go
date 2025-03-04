package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/trisatya23/go-restful-api-master/app"
	"github.com/trisatya23/go-restful-api-master/controller"
	"github.com/trisatya23/go-restful-api-master/helper"
	"github.com/trisatya23/go-restful-api-master/model/domain"
	"github.com/trisatya23/go-restful-api-master/repository"
	"github.com/trisatya23/go-restful-api-master/service"
	"log"
)

func main() {

	server := fiber.New()

	// Initialize Database
	db := app.NewDB()

	// Run Auto Migration (Opsional, bisa dihapus jika tidak diperlukan)
	err := db.AutoMigrate(&domain.Category{})
	helper.PanicIfError(err)

	// Initialize Validator
	validate := validator.New()

	// Initialize Repository, Service, and Controller
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// Setup Routes
	app.NewRouter(server, categoryController)

	// Start Server
	log.Println("Server running on port 8080")
	err = server.Listen(":8080")
	helper.PanicIfError(err)
}
