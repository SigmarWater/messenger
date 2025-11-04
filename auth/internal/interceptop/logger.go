package interceptop

import (
	"context"
	"path"
	"log"
	"time"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
) 


func LoggerInceptop() grpc.UnaryServerInterceptor{
	return func(
		ctx context.Context, 
		req interface{}, 
		info *grpc.UnaryServerInfo, 
		handler grpc.UnaryHandler,
	)(interface{}, error){
		// Извлекаем начало вызова метода
		method := path.Base(info.FullMethod)

		// Логируем начало вызова метода
		log.Printf("Started gRPC method %s\n", method)

		// Засекаем время начала выполнения
		startTime := time.Now() 

		// Вызов обработчика
		resp, err := handler(ctx, req)

		// Вычисляем длительность выполнения 
		duration := time.Since(startTime)

		// Форматируем сообщение в зависимости от результата
		if err != nil{
			st, _ := status.FromError(err)
			log.Printf("Failed finished grpc method %s with code %s: %v (took: %v)\n", method, st.Code(), err, duration)
		}else{
			log.Printf("Success Finished grpc method %s", method)
		}

		return resp, err
	}
}