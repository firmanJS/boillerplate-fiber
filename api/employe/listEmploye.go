package employe

import (
	. "github.com/firmanJS/boillerplate-fiber/config"
	. "github.com/firmanJS/boillerplate-fiber/models"
	"github.com/firmanJS/boillerplate-fiber/helpers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAll(ctx *fiber.Ctx) error {
	query := bson.D{{}}
	cursor, queryError := Instance.Database.Collection("employe").Find(ctx.Context(), query)
	if queryError != nil {
		return helpers.CrudResponse(ctx, "Get", queryError)
	}

	var employe []Employe = make([]Employe, 0)

	// iterate the cursor and decode each item into a Todo
	if err := cursor.All(ctx.Context(), &employe); err != nil {
		return helpers.MsgResponse(ctx, "get data unsuccesfully", err)
	}

	return helpers.CrudResponse(ctx, "Get", employe)
}
