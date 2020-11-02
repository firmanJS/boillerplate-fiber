package routers

import (
	"github.com/firmanJS/boillerplate-fiber/api/role"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes Role
func RoleRoutes(api fiber.Router) {

	route := api.Group("/role")
	route.Get("/", role.GetAll)

}
