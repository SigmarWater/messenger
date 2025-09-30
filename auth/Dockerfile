FROM golang:1.25-alpine3.22  AS builder 
WORKDIR /app 
COPY go.mod go.sum ./
RUN go mod download  
COPY ./cmd/userApi/main.go . 
RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o auth_service main.go 

FROM alpine:3.22.1   
WORKDIR /app
COPY --from=builder app/auth_service ./
CMD ["./auth_service"] 