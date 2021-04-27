package employe

import (
	"github.com/firmanJS/boillerplate-fiber/config"
	"github.com/firmanJS/boillerplate-fiber/helpers"
	"github.com/firmanJS/boillerplate-fiber/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSingle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	employeId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	collection := config.Instance.Database.Collection("employe")

	query := bson.D{{Key: "_id", Value: employeId}}
	rawRecord := collection.FindOne(ctx.Context(), query)
	record := &models.Employe{}
	rawRecord.Decode(record)

	if rawRecord.Err() != nil {
		return helpers.NotFoundResponse(ctx, "Data not found in database")
	} else {
		return helpers.CrudResponse(ctx, "Get", record)
	}
}
