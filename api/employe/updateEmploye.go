package employe

import (
	"time"

	. "github.com/firmanJS/boillerplate-fiber/config"
	"github.com/firmanJS/boillerplate-fiber/helpers"
	. "github.com/firmanJS/boillerplate-fiber/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateSingle(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	employe := new(Employe)
	employeId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return helpers.BadResponse(ctx, "Bad Request", parseError.Error())
	}

	parsingError := ctx.BodyParser(employe)
	if parsingError != nil {
		helpers.ServerResponse(ctx, parsingError.Error(), parsingError.Error())
	}

	collection := Instance.Database.Collection("employe")

	// check if the record is there
	query := bson.D{{Key: "_id", Value: employeId}}
	rawRecord := collection.FindOne(ctx.Context(), query)
	record := &Employe{}
	rawRecord.Decode(record)

	// update the record
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employe.Name},
				{Key: "salary", Value: employe.Salary},
				{Key: "age", Value: employe.Age},
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
