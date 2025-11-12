# 📋 Guia de Imports - MongoDB Driver v2

## ⚠️ IMPORTANTE: Versão do Driver

Este projeto usa **MongoDB Driver v2** (`go.mongodb.org/mongo-driver/v2`).

**NUNCA** misture com a versão v1 (`go.mongodb.org/mongo-driver`).

---

## ✅ IMPORTS CORRETOS

### Para trabalhar com BSON:
```go
import "go.mongodb.org/mongo-driver/v2/bson"

// Usar:
type ID = bson.ObjectID
id, err := bson.ObjectIDFromHex(hexString)
filter := bson.D{{Key: "_id", Value: id}}
```

### Para trabalhar com tipos de configuração do projeto:
```go
import rest_err "github.com/dlima78/gocourse/src/configuration"

// Usar tipos padrão do projeto:
id, err := rest_err.NewObjectIDFromHex(hexString)
isValid := rest_err.IsValidObjectID(hexString)
newID := rest_err.NewObjectID()
```

### Para conectar ao MongoDB:
```go
import "go.mongodb.org/mongo-driver/v2/mongo"
import "go.mongodb.org/mongo-driver/v2/mongo/options"
```

---

## ❌ IMPORTS INCORRETOS (NUNCA USE)

```go
// ❌ ERRADO - Versão v1 (deprecated)
import "go.mongodb.org/mongo-driver/bson/primitive"
type ID = primitive.ObjectID
id, _ := primitive.ObjectIDFromHex(hexString)

// ❌ ERRADO - Versão antiga do mongo-driver
import "go.mongodb.org/mongo-driver/bson"
```

---

## 📝 Template para Repository

```go
package repository

import (
	"context"
	"fmt"
	"os"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByID(id string) error {
	// ✅ Usar rest_err.NewObjectIDFromHex
	objectId, err := rest_err.NewObjectIDFromHex(id)
	if err != nil {
		logger.Error("Invalid ID format", err, zap.String("id", id))
		return rest_err.NewBadRequestError("Invalid ID format")
	}

	// ✅ Usar bson.D para filtros
	filter := bson.D{{Key: "_id", Value: objectId}}
	
	collection := ur.databaseConnection.Collection("users")
	result := collection.FindOne(context.Background(), filter)
	
	return result.Err()
}
```

---

## 🔍 Verificar Imports

Execute o comando para verificar se há imports incorretos:

```bash
# Procurar por imports v1
grep -r "mongo-driver/bson/primitive" src/
grep -r "mongo-driver/bson[^/]" src/

# Procurar por imports v1 (alternativo)
grep -r "primitive\." src/ | grep -v "go.mongodb.org/mongo-driver/v2"
```

Se encontrar algo, corrija imediatamente!

---

## 📦 Verificar Dependências

```bash
# Listar todas as versões do mongo-driver
go list -m all | grep mongo-driver

# Verificar se há duplicatas (deve mostrar apenas v2)
go mod graph | grep mongo-driver

# Limpar dependências
go mod tidy
```

---

## 🚀 Checklist para Novos Arquivos

Ao criar um novo arquivo de repository, verifique:

- [ ] Imports: `go.mongodb.org/mongo-driver/v2/bson`
- [ ] Imports: `go.mongodb.org/mongo-driver/v2/mongo`
- [ ] Usar `rest_err.NewObjectIDFromHex()` para converter IDs
- [ ] Usar `bson.D` para filtros
- [ ] Usar `bson.ObjectID` para tipos (via `rest_err.ObjectID`)
- [ ] Não há imports de `primitive`
- [ ] Não há imports antigos de `mongo-driver/bson`

---

## 💡 Dúvidas?

Se tiver dúvida sobre qual import usar, busque por exemplos nos arquivos:
- `src/model/repository/find_user_repository.go` ✅
- `src/model/repository/create_user_repository.go` ✅
- `src/model/repository/entity/user_entity.go` ✅
