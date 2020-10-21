package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employe struct {
	ID     primitive.ObjectID  `json:"_id" bson:"_id,omitempty"`
	Name   string  `json:"name,omitempty" bson:"name,omitempty""`
	Salary float64 `json:"salary,omitempty" bson:"salary,omitempty"`
	Age    int64 `json:"age,omitempty" bson:"age,omitempty`
	CreatedAt time.Time  `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt time.Time  `json:"updatedAt" bson:"updatedAt" binding:"required"`
}
