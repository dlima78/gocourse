FROM golang:1.25.1 AS builder 

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gocourse .

FROM alpine:latest AS runner

RUN adduser -D dulima

WORKDIR /app

COPY --from=builder /app/gocourse .
COPY --from=builder /app/.env .

RUN chown -R dulima:dulima /app
RUN chmod +x /app/gocourse

USER dulima

EXPOSE 8080

CMD [ "./gocourse" ]