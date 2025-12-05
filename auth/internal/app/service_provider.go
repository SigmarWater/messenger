package app

import (
	"context"
	"github.com/SigmarWater/messenger/auth/internal/closer"
	"github.com/SigmarWater/messenger/auth/internal/repository/auth_cache"
	redigo "github.com/gomodule/redigo/redis"
	"log"

	"github.com/SigmarWater/messenger/auth/internal/api/auth"
	"github.com/SigmarWater/messenger/auth/internal/config"
	"github.com/SigmarWater/messenger/auth/internal/config/env"
	"github.com/SigmarWater/messenger/auth/internal/repository"
	"github.com/SigmarWater/messenger/auth/internal/repository/users"
	"github.com/SigmarWater/messenger/auth/internal/service"
	serviceUser "github.com/SigmarWater/messenger/auth/internal/service/users"
	"github.com/SigmarWater/messenger/platform/pkg/cache"
	"github.com/SigmarWater/messenger/platform/pkg/cache/redis"
	"github.com/jackc/pgx/v4/pgxpool"
)

type serviceProvider struct {
	pgConfig    config.PGConfig
	grpcConfig  config.GRPCConfig
	httpConfig  config.HTTPConfig
	redisConfig config.CacheConfig

	pgPool              *pgxpool.Pool
	userRepository      repository.UserRepository
	userCacheRepository repository.UserCache

	redisPool   *redigo.Pool
	redisClient cache.RedisClient

	serviceUser service.UsersService

	userImpl *auth.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPgConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s\n", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		grpcCgf, err := env.NewGrpcConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v\n", err.Error())
		}

		s.grpcConfig = grpcCgf
	}
	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		httpCfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v\n", err.Error())
		}

		s.httpConfig = httpCfg
	}
	return s.httpConfig
}

func (s *serviceProvider) RedisConfig() config.CacheConfig {
	if s.redisConfig == nil {
		redisCfg, err := env.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to get redis config: %v\n", err.Error())
		}

		s.redisConfig = redisCfg
	}
	return s.redisConfig
}

func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to db: %v\n", err)
		}

		if err := pool.Ping(ctx); err != nil {
			log.Fatalf("failed to ping to db: %v\n", err)
		}

		closer.Add(func() error {
			pool.Close()
			return nil
		})
		s.pgPool = pool
	}

	return s.pgPool
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = users.NewPostgresUserRepository(s.PgPool(ctx))
	}
	return s.userRepository
}

func (s *serviceProvider) CachePool(ctx context.Context) *redigo.Pool {
	if s.redisPool == nil {
		_ = s.RedisConfig()
		s.redisPool = &redigo.Pool{
			MaxIdle:     s.redisConfig.MaxIdle(),
			IdleTimeout: s.redisConfig.IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", s.redisConfig.Address())
			},
		}
	}
	return s.redisPool
}

func (s *serviceProvider) RedisClient(ctx context.Context) cache.RedisClient {
	if s.redisClient == nil {
		s.redisClient = redis.NewClient(s.CachePool(ctx), s.RedisConfig().ConnectionTimeout())
	}
	return s.redisClient
}

func (s *serviceProvider) UserCache(ctx context.Context) repository.UserCache {
	if s.userCacheRepository == nil {
		s.userCacheRepository = auth_cache.NewRepository(s.RedisClient(ctx))
	}
	return s.userCacheRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UsersService {
	if s.serviceUser == nil {
		s.serviceUser = serviceUser.NewService(s.UserRepository(ctx), s.UserCache(ctx), s.RedisConfig().CacheTTL())
	}

	return s.serviceUser
}

func (s *serviceProvider) UserImpl(ctx context.Context) *auth.Implementation {
	if s.userImpl == nil {
		s.userImpl = auth.NewImplementation(s.UserService(ctx))
	}
	return s.userImpl
}
