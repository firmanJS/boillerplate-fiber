package service

import (
	"github.com/asaskevich/govalidator"
	"github.com/firmanJS/boillerplate-fiber/config"
	"github.com/firmanJS/boillerplate-fiber/helpers"
	"github.com/firmanJS/boillerplate-fiber/models"
	"github.com/firmanJS/boillerplate-fiber/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateNew(ctx *fiber.Ctx) error {

	collection := config.Instance.Database.Collection("employe")

	// create a new record
	employe := new(models.Employe)
	employe.CreatedAt = utils.MakeTimestamp()
	employe.UpdatedAt = utils.MakeTimestamp()

	if errors := ctx.BodyParser(employe); errors != nil {
		_, err := govalidator.ValidateStruct(employe)

		if err != nil {
			return helpers.ServerResponse(ctx, err.Error(), err)
		}

		return helpers.ServerResponse(ctx, errors.Error(), errors)
	} else {
		if result, errs := collection.InsertOne(ctx.Context(), employe); errs != nil {
			return helpers.ServerResponse(ctx, errs.Error(), errs.Error())
		} else {
			filter := bson.D{{Key: "_id", Value: result.InsertedID}}
			createdRecord := collection.FindOne(ctx.Context(), filter)
			createdemploye := &models.Employe{}
			createdRecord.Decode(createdemploye)

			return helpers.CrudResponse(ctx, "Create", createdemploye)
		}
	}
}
