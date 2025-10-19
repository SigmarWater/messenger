FROM golang:1.25-alpine3.22  AS builder 
WORKDIR /app 
COPY go.mod go.sum /app/
RUN go mod download  
COPY . /app/ 
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o auth_service ./cmd/userApi/main.go 

FROM alpine:3.22.1   
WORKDIR /app
COPY --from=builder /app/auth_service /app/
COPY --from=builder /app/../postgres /app/postgres
CMD ["./auth_service"] 
