package role

import (
	"time"
	. "github.com/firmanJS/boillerplate-fiber/config"
	"github.com/firmanJS/boillerplate-fiber/helpers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateSingle(ctx *fiber.Ctx) error {
	// check data
	id := ctx.Params("id")
	role := new(Role)
	roleId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	parsingError := ctx.BodyParser(role)
	if parsingError != nil {
		helpers.ServerResponse(ctx, parsingError.Error(), parsingError.Error())
	}

	collection := Instance.Database.Collection("role")

	// check if the record is there
	query := bson.D{{Key: "_id", Value: roleId}}
	rawRecord := collection.FindOne(ctx.Context(), query)
	record := &Role{}
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

