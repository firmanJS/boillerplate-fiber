package config

import (
	"context"
	"time"

	"github.com/firmanJS/fiber-with-mongo/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var Instance MongoInstance

// Create database connection
func Connect() error {
	DatabaseConnection := config.Config("MONGO_HOST")
	DatabaseName := config.Config("MONGO_DB_NAME")
	client, err := mongo.NewClient(options.Client().ApplyURI(DatabaseConnection))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(DatabaseName)

	if err != nil {
		return err
	}

	Instance = MongoInstance{
		Client:   client,
		Database: db,
	}

	return nil
}
