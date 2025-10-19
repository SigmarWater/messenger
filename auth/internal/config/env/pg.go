package env

import (
	"errors"
	"os"

	"github.com/SigmarWater/messenger/auth/internal/config"
)

var _ config.PGConfig = (*pgConfig)(nil)

const (
	dsnEnvName = "PG_DNS"
)

type pgConfig struct {
	dns string
}

func NewPgConfig() (*pgConfig, error) {
	dns := os.Getenv(dsnEnvName)
	if len(dns) == 0 {
		return nil, errors.New("not found such value")
	}
	return &pgConfig{
		dns: dns,
	}, nil
}

func (c *pgConfig) DSN() string {
	return c.dns
}
