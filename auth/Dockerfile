FROM golang:1.25-alpine3.22  AS builder 
WORKDIR /app 
# Копируем go.mod и go.sum для platform (на уровень выше, как указано в replace)
COPY platform/go.mod platform/go.sum /platform/
# Копируем go.mod и go.sum для auth
COPY auth/go.mod auth/go.sum /app/
# Скачиваем зависимости (platform уже доступен через replace)
RUN go mod download  
# Копируем исходный код platform
COPY platform/ /platform/
# Копируем исходный код auth
COPY auth/ /app/ 
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o auth_service ./cmd/userApi/main.go 

FROM alpine:3.22.1   
WORKDIR /app
COPY --from=builder /app/auth_service /app/
COPY postgres/migrations/.env /app/.env
EXPOSE 8083 50000
CMD ["./auth_service"] 
