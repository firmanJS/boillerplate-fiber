package role

import (
	. "github.com/firmanJS/boillerplate-fiber/config"
	. "github.com/firmanJS/boillerplate-fiber/models"
	"github.com/firmanJS/boillerplate-fiber/utils"
	"github.com/firmanJS/boillerplate-fiber/helpers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/asaskevich/govalidator"
)

func CreateNew(ctx *fiber.Ctx) error {
	
	collection := Instance.Database.Collection("role")

	// create a new record
	role := new(Role)
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
			createdRole := &Role{}
			createdRecord.Decode(createdRole)

			return helpers.CrudResponse(ctx, "Create", createdRole)
		}
	}
}
