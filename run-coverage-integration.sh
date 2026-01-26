#!/bin/bash
set -e

# Build do projeto
echo "Compilando projeto..."
go build ./src/...

# Executar testes de integração do repository com cobertura
echo "Executando testes de integração com cobertura..."
go test ./src/model/repository/... \
  -coverprofile=coverage_integration.out \
  -covermode=atomic \
  -timeout=120s \
  -v

# Gerar relatório HTML
echo "Gerando relatório HTML..."
go tool cover -html=coverage_integration.out -o coverage_integration.html

# Exibir resumo de cobertura
echo ""
echo "=== Resumo de Cobertura (Testes de Integração) ==="
go tool cover -func=coverage_integration.out | tail -5

echo ""
echo "Relatório HTML gerado: coverage_integration.html"
