package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trisatya23/go-restful-api-master/controller"
	"github.com/trisatya23/go-restful-api-master/middleware"
)

func NewRouter(app *fiber.App, categoryController controller.CategoryController) {
	authMiddleware := middleware.NewAuthMiddleware()

	api := app.Group("/api", authMiddleware)
	categories := api.Group("/categories")

	categories.Get("/", categoryController.FindAll)
	categories.Get("/:categoryId", categoryController.FindById)
	categories.Post("/", categoryController.Create)
	categories.Put("/:categoryId", categoryController.Update)
	categories.Delete("/:categoryId", categoryController.Delete)
}
