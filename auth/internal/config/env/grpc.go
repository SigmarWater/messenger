package env

import (
	"errors"
	"github.com/SigmarWater/messenger/auth/internal/config"
	"net"
	"os"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

const (
	grpcEnvHost = "GRPC_HOST"
	grpcEnvPort = "GRPC_PORT_AUTH"
)

type grpcConfig struct {
	host string
	port string
}

func NewGrpcConfig() (*grpcConfig, error) {
	host := os.Getenv(grpcEnvHost)
	if len(host) == 0 {
		return nil, errors.New("not found host key")
	}

	port := os.Getenv(grpcEnvPort)
	if len(port) == 0 {
		return nil, errors.New("not found port key")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (c *grpcConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
