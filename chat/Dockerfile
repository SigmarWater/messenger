FROM golang:1.25-alpine3.22 AS builder 
WORKDIR /app 
COPY chat/go.mod chat/go.sum /app/ 
RUN go mod download  
COPY chat/ /app/ 
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o chat_server ./cmd/chat_server/main.go 

FROM alpine:3.22.1 
WORKDIR /app 
COPY --from=builder /app/chat_server /app/
COPY postgres/ /app/postgres/
CMD ["./chat_server"] 





