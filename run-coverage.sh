go test ./src/... -coverprofile=coverage.out -v
go tool cover -html=coverage.out -o coverage.html
start chrome coverage.html