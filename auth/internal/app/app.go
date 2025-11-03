package app

import (
	"context"
	"log"
	"net"
	"sync"


	"net/http"
	

	"github.com/SigmarWater/messenger/auth/internal/closer"
	"github.com/SigmarWater/messenger/auth/internal/config"
	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)


type app struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer *http.Server 
}

func NewApp(ctx context.Context) (*app, error) {
	a := &app{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *app) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := &sync.WaitGroup{} 

	wg.Add(2)

	go func(){
		defer wg.Done() 

		err := a.runGRPCServer()

		if err != nil{
			log.Fatalf("failed to run GRPC server")
		}
	}()

	go func(){
		defer wg.Done() 

		err := a.runHTTPServer()

		if err != nil{
			log.Fatalf("failed to run GRPC server")
		}
	}()

	wg.Wait() 

	return nil
}


func (a *app) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *app) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *app) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *app) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	pb.RegisterUserAPIServer(a.grpcServer, a.serviceProvider.UserImpl(ctx))

	return nil
}

func (a *app) initHTTPServer(ctx context.Context) error{
	// Создаем мультиплексор для HTTP запросов
	mux := runtime.NewServeMux()

	// Настраиваем опции для соединения с gRPC сервером
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Регистрируем gRPC-gateway хендлеры
	err := pb.RegisterUserAPIHandlerFromEndpoint(
		ctx,
		mux,
		a.serviceProvider.grpcConfig.Address(),
		opts,
	)

	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
		return err 
	}

	a.httpServer = &http.Server{
		Addr: a.serviceProvider.HTTPConfig().Address() ,
		Handler: mux,
	}

	return nil
}

func (a *app) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (a *app) runHTTPServer() error{
	log.Printf("HTTP server is running on %s", a.serviceProvider.httpConfig.Address())

	err := a.httpServer.ListenAndServe()

	if err != nil{
		return err 
	}

	return nil
}
