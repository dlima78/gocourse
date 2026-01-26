#!/bin/bash
set -e

# Build do projeto
echo "Compilando projeto..."
go build ./src/...

# Executar apenas testes unitários com cobertura
echo "Executando testes unitários com cobertura..."
go test ./src/... \
  -short \
  -coverprofile=coverage.out \
  -covermode=atomic \
  -timeout=60s \
  -v

# Gerar relatório HTML
echo "Gerando relatório HTML..."
go tool cover -html=coverage.out -o coverage.html

# Exibir resumo de cobertura
echo ""
echo "=== Resumo de Cobertura (Testes Unitários) ==="
go tool cover -func=coverage.out | tail -5

echo ""
echo "Relatório HTML gerado: coverage.html"
echo ""
echo "Para rodar testes de integração com cobertura:"
echo "go test ./src/model/repository/... -coverprofile=coverage_integration.out -v"
