package routers

import (
	api "github.com/firmanJS/boillerplate-fiber/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	apiRoutes := app.Group("/api", logger.New())

	apiRoutes.Get("/", api.Index)
	EmployeRoutes(apiRoutes)
	RoleRoutes(apiRoutes)
}
