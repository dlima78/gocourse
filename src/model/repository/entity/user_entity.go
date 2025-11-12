package entity

import "go.mongodb.org/mongo-driver/v2/bson"

// UserEntity representa um documento de usuário no MongoDB
//
// ⚠️  IMPORTANTE - Versão do Driver:
// Este struct usa bson.ObjectID do mongo-driver v2
// Nunca misture com primitive.ObjectID da v1
//
// Para converter ID de string para ObjectID:
//
//	import rest_err "github.com/dlima78/gocourse/src/configuration"
//	oid, err := rest_err.NewObjectIDFromHex(hexString)
type UserEntity struct {
	// ID é o ObjectID primário do MongoDB (v2)
	// Tag json:"_id" para compatibilidade com JSON/REST
	// Tag bson:"_id,omitempty" para armazenamento no MongoDB
	ID       bson.ObjectID `json:"_id" bson:"_id,omitempty"`
	Email    string        `bson:"email,omitempty"`
	Password string        `bson:"password,omitempty"`
	Name     string        `bson:"name,omitempty"`
	Age      int8          `bson:"age,omitempty"`
}
