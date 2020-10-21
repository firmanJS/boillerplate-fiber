package employe

import (
	. "github.com/firmanJS/boillerplate-fiber/config"
	"github.com/firmanJS/boillerplate-fiber/helpers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteSingle(ctx *fiber.Ctx) error {
	// check data
	id := ctx.Params("id")

	employeId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	// get collection
	collection := Instance.Database.Collection("employe")

	// check if the record is there
	query := bson.D{{Key: "_id", Value: employeId}}
	result, deleteError := collection.DeleteOne(ctx.Context(), &query)

	if deleteError != nil {
		return helpers.ServerResponse(ctx, deleteError.Error(), deleteError.Error())
	}

	// check if item was deleted
	if result.DeletedCount < 1 {
		return helpers.NotFoundResponse(ctx, "Data not found in database")
	} else {
		return helpers.CrudResponse(ctx, "Deleted", result)
	}
}