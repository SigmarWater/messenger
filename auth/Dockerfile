FROM golang:1.25-alpine3.22 AS builder 
WORKDIR /app 
COPY ../go.mod ../go.sum /app/ 
RUN go mod download  
COPY ./cmd/userApi/main.go /app/ 
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o user_api ./main.go 

FROM alpine:3.22.1 
WORKDIR /app 
COPY --from=builder /app/user_api /app/  
CMD ["./user_api"] 

