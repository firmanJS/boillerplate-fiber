package routers

import (
	"github.com/firmanJS/boillerplate-fiber/api/employe"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes Employe
func EmployeRoutes(api fiber.Router) {

	route := api.Group("/employe")
	route.Get("/", employe.GetAll)
	route.Post("/", employe.CreateNew)
	route.Delete("/:id", employe.DeleteSingle)
	route.Get("/:id", employe.GetSingle)
	route.Put("/:id", employe.UpdateSingle)

}
