package env

import(
	"os"
	"github.com/SigmarWater/messenger/auth/internal/config"
	"errors"
)

var _ config.PGConfig = (*pgConfig)(nil)


const(
	dsnEnvName = "PG_DNS"
)

type pgConfig struct{
	dns string 
}

func NewPgConfig() (*pgConfig, error){
	dns := os.Getenv(dsnEnvName) 
	if len(dns) == 0{
		return nil, errors.New("not found such value")
	}
	return &pgConfig{
		dns: dns,
	}, nil 
}

func (c *pgConfig) DNS() string{
	return c.dns
}