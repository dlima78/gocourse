package rest_err

import "go.mongodb.org/mongo-driver/v2/bson"

// ObjectID é o tipo padrão para IDs do MongoDB
//
// IMPORTANTE: Sempre use este tipo ao invés de primitive.ObjectID
// Use bson.ObjectIDFromHex(hexString) para converter de string
type ObjectID = bson.ObjectID

// D é o tipo padrão para documentos BSON
type D = bson.D

// E é o tipo padrão para elementos BSON
type E = bson.E

// NewObjectIDFromHex converte uma string hex para ObjectID
//
// Exemplo:
//
//	oid, err := NewObjectIDFromHex("507f1f77bcf86cd799439011")
func NewObjectIDFromHex(hex string) (ObjectID, error) {
	return bson.ObjectIDFromHex(hex)
}

// NewObjectID gera um novo ObjectID único
func NewObjectID() ObjectID {
	return bson.NewObjectID()
}

// IsValidObjectID verifica se uma string é um ObjectID válido
func IsValidObjectID(hex string) bool {
	_, err := bson.ObjectIDFromHex(hex)
	return err == nil
}
