package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/firmanJS/boillerplate-fiber/api"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	apiRoutes := app.Group("/api", logger.New())
	
	apiRoutes.Get("/", api.Index)
	EmployeRoutes(apiRoutes)
}