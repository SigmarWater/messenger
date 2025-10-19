package app

import (
	"context"
	"log"
	"net"

	"github.com/SigmarWater/messenger/auth/internal/closer"
	"github.com/SigmarWater/messenger/auth/internal/config"
	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type app struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
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

	return a.runGRPCServer()
}

func (a *app) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
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
	err := config.Load("./postgres/migrations/.env")
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
