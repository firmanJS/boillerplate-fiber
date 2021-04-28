package role

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

	collection := config.Instance.Database.Collection("role")

	// create a new record
	role := new(models.Role)
	role.CreatedAt = utils.MakeTimestamp()
	role.UpdatedAt = utils.MakeTimestamp()

	if errors := ctx.BodyParser(role); errors != nil {
		_, err := govalidator.ValidateStruct(role)

		if err != nil {
			return helpers.ServerResponse(ctx, err.Error(), err)
		}

		return helpers.ServerResponse(ctx, errors.Error(), errors)
	} else {
		if result, errs := collection.InsertOne(ctx.Context(), role); errs != nil {
			return helpers.ServerResponse(ctx, errs.Error(), errs.Error())
		} else {
			filter := bson.D{{Key: "_id", Value: result.InsertedID}}
			createdRecord := collection.FindOne(ctx.Context(), filter)
			createdRole := &models.Role{}
			createdRecord.Decode(createdRole)

			return helpers.CrudResponse(ctx, "Create", createdRole)
		}
	}
}
