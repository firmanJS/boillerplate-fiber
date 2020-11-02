package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// perbedaan menggunakan omitempty dan tidak jika menggunakan 
// omitempty ketika post json kosong tidak terinsert ke document
// jika menggunkana emitpty nanti data akan duplikat juga pas insert

type Role struct {
	ID     primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	RoleName   string   `json:"roleName" valid:"required~roleName is blank" bson:"roleName"`
	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" bson:"updatedAt"`
}
