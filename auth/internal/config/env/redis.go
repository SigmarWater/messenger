package env

import (
	"errors"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	redisHostEnvName              = "REDIS_HOST"
	redisPortEnvName              = "REDIS_PORT"
	redisConnectionTimeoutEnvName = "REDIS_CONNECTION_TIMEOUT"
	redisMaxIdleEnvName           = "REDIS_MAX_IDLE"
	redisIdleTimeoutEnvName       = "REDIS_IDLE_TIMEOUT"
	redisCacheTTL                 = "REDIS_CACHE_TTL"
)

type redisConfig struct {
	host              string
	port              string
	connectionTimeout time.Duration
	maxIdle           int
	idleTimeout       time.Duration
	cacheTTL          time.Duration
}

func NewRedisConfig() (*redisConfig, error) {
	host := os.Getenv(redisHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("not found host key")
	}
	port := os.Getenv(redisPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("not found port key")
	}
	connectionTimeout := os.Getenv(redisConnectionTimeoutEnvName)
	if len(connectionTimeout) == 0 {
		return nil, errors.New("not found ConnectionTimeout key")
	}

	connectionTimeoutParsed, err := time.ParseDuration(connectionTimeout)
	if err != nil {
		return nil, errors.New("error parse connection time")
	}

	maxIdle := os.Getenv(redisMaxIdleEnvName)
	if len(maxIdle) == 0 {
		return nil, errors.New("not fount MaxIdleEnvName key")
	}

	maxIdleParsed, err := strconv.ParseInt(maxIdle, 10, 64)
	if err != nil {
		return nil, errors.New("error parse maxIdle")
	}

	idleTimeout := os.Getenv(redisIdleTimeoutEnvName)
	if len(idleTimeout) == 0 {
		return nil, errors.New("not found IdleTimeout key")
	}

	idleTimeoutParsed, err := time.ParseDuration(idleTimeout)
	if err != nil {
		return nil, errors.New("error parse idleTimeout time")
	}

	cacheTTL := os.Getenv(redisCacheTTL)
	if len(cacheTTL) == 0 {
		return nil, errors.New("not found cacheTTL key")
	}

	cacheTTLParsed, err := time.ParseDuration(cacheTTL)
	if err != nil {
		return nil, errors.New("error parse cacheTtl time")
	}

	return &redisConfig{
		host:              host,
		port:              port,
		connectionTimeout: connectionTimeoutParsed,
		maxIdle:           int(maxIdleParsed),
		idleTimeout:       idleTimeoutParsed,
		cacheTTL:          cacheTTLParsed,
	}, nil
}

func (r *redisConfig) Address() string {
	return net.JoinHostPort(r.host, r.port)
}

func (r *redisConfig) ConnectionTimeout() time.Duration {
	return r.connectionTimeout
}

func (r *redisConfig) MaxIdle() int {
	return r.maxIdle
}

func (r *redisConfig) IdleTimeout() time.Duration {
	return r.idleTimeout
}

func (r *redisConfig) CacheTTL() time.Duration {
	return r.cacheTTL
}
