package role

import (
	"github.com/firmanJS/boillerplate-fiber/config"
	"github.com/firmanJS/boillerplate-fiber/helpers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteSingle(ctx *fiber.Ctx) error {
	// check data
	id := ctx.Params("id")

	roleId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	// get collection
	collection := config.Instance.Database.Collection("role")

	// check if the record is there
	query := bson.D{{Key: "_id", Value: roleId}}
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

func DeleteAll(ctx *fiber.Ctx) error {
	// get collection
	collection := config.Instance.Database.Collection("role")

	// check if the record is there
	deleteResult := collection.Drop(ctx.Context())

	return helpers.CrudResponse(ctx, "Deleted", deleteResult)
}
