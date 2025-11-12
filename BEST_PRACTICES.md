# 🚀 Melhores Práticas do Projeto

## MongoDB Driver - v2 SOMENTE

Este projeto usa **exclusivamente** `go.mongodb.org/mongo-driver/v2`.

### ❌ Erro Comum que Cometemos

Misturar versões do driver:
```go
// ❌ ERRADO
import "go.mongodb.org/mongo-driver/bson/primitive"  // v1
import "go.mongodb.org/mongo-driver/v2/bson"         // v2

// Isso causa erro: tipos incompatíveis!
var oid primitive.ObjectID  // v1
filter := bson.D{{...}}     // v2 - NÃO FUNCIONA COM v1!
```

### ✅ Forma Correta

Usar APENAS tipos do v2:
```go
import "go.mongodb.org/mongo-driver/v2/bson"

var oid bson.ObjectID = bson.NewObjectID()
filter := bson.D{{Key: "_id", Value: oid}}  // FUNCIONA!
```

---

## 🔧 Como Verificar e Corrigir

### 1. Verificar dependências
```bash
go mod graph | grep mongo-driver  # Deve mostrar APENAS v2
```

### 2. Procurar imports incorretos
```bash
grep -r "bson/primitive" src/     # Deve estar vazio
grep -r "mongo-driver[^/v]" src/  # Deve estar vazio
```

### 3. Limpar dependências
```bash
go mod tidy
go mod verify
```

### 4. Validar com linter
```bash
golangci-lint run ./...
```

---

## 📁 Estrutura de Arquivos Importantes

```
.golangci.yml                          # Configuração de linter
MONGODB_IMPORTS_GUIDE.md              # Guia detalhado de imports
src/
├── configuration/
│   ├── mongo_types.go                # Tipos padrão do projeto ✅
│   └── rest_err.go                   
├── model/
│   └── repository/
│       ├── find_user_repository.go   # Exemplo: imports corretos ✅
│       ├── create_user_repository.go # Exemplo: imports corretos ✅
│       └── entity/
│           └── user_entity.go        # Exemplo: struct com bson.ObjectID ✅
```

---

## 🛡️ Proteção contra Regressão

### Antes de fazer commit:
```bash
# 1. Verificar imports
grep -r "bson/primitive" src/ && echo "❌ ENCONTRADO IMPORT v1" || echo "✅ Sem imports v1"

# 2. Limpar dependências
go mod tidy
go mod verify

# 3. Compilar
go build -o tmp/main .

# 4. Rodar linter
golangci-lint run ./...

# 5. Se tudo ok, fazer commit
git add .
git commit -m "fix: descrição da mudança"
```

### Adicionar pre-commit hook (opcional):
Crie `.git/hooks/pre-commit`:
```bash
#!/bin/bash
grep -r "bson/primitive" src/ && {
  echo "❌ ERRO: Encontrado import v1 do mongo-driver!"
  echo "Use 'go.mongodb.org/mongo-driver/v2/bson' ao invés!"
  exit 1
}
go mod verify || exit 1
go build -o tmp/main . || exit 1
```

---

## 📚 Referências Úteis

1. **mongo-driver v2**: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2
2. **Guia de imports**: Ver `MONGODB_IMPORTS_GUIDE.md`
3. **Tipos padrão**: Ver `src/configuration/mongo_types.go`
4. **Exemplos de uso**: Ver `src/model/repository/*.go`

---

## ✨ Resumo

| Antes ❌ | Depois ✅ |
|---------|---------|
| `primitive.ObjectID` | `bson.ObjectID` (v2) |
| `primitive.ObjectIDFromHex()` | `bson.ObjectIDFromHex()` ou `rest_err.NewObjectIDFromHex()` |
| Misturar v1 e v2 | Usar APENAS v2 |
| Sem verificação | `go mod verify` + `golangci-lint` |
| Sem documentação | `MONGODB_IMPORTS_GUIDE.md` + comentários |

---

*Última atualização: 12/11/2025 - Corrigido erro de mistura de versões do mongo-driver*
