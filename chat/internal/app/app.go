package app

import (
	"context"
	"log"
	"net"

	"github.com/SigmarWater/messenger/chat/internal/closer"
	"github.com/SigmarWater/messenger/chat/internal/config"
	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.initRunGRPCServer(ctx)
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGrpcServer,
		a.initRunGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load("./.env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	pb.RegisterChatApiServer(a.grpcServer, a.serviceProvider.Impl(ctx))
	reflection.Register(a.grpcServer)
	return nil
}

func (a *App) initRunGRPCServer(_ context.Context) error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GrpcConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GrpcConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
