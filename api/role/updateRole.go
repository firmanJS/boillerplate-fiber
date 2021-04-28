package role

import (
	"time"

	"github.com/firmanJS/boillerplate-fiber/config"
	"github.com/firmanJS/boillerplate-fiber/helpers"
	"github.com/firmanJS/boillerplate-fiber/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateSingle(ctx *fiber.Ctx) error {
	// check data
	id := ctx.Params("id")
	role := new(models.Role)
	roleId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	parsingError := ctx.BodyParser(role)
	if parsingError != nil {
		helpers.ServerResponse(ctx, parsingError.Error(), parsingError.Error())
	}

	collection := config.Instance.Database.Collection("role")

	// check if the record is there
	query := bson.D{{Key: "_id", Value: roleId}}
	rawRecord := collection.FindOne(ctx.Context(), query)
	record := &models.Role{}
	rawRecord.Decode(record)

	if rawRecord.Err() != nil {
		return helpers.NotFoundResponse(ctx, "Data not found in database")
	}

	// update the record
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "rolename", Value: role.RoleName},
				{Key: "updatedAt", Value: time.Now()},
			},
		},
	}
	result, updateError := collection.UpdateOne(ctx.Context(), query, update)
	if updateError != nil {
		return helpers.ServerResponse(ctx, updateError.Error(), updateError.Error())
	}

	return helpers.CrudResponse(ctx, "Update", result)
}
