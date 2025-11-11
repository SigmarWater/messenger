package auth_cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/SigmarWater/messenger/auth/internal/model"
	"github.com/SigmarWater/messenger/auth/internal/repository/converter"
	cacheModel "github.com/SigmarWater/messenger/auth/internal/repository/model"
	"github.com/SigmarWater/messenger/platform/pkg/cache"
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

const (
	cacheKeyPrefix = "auth:user:"
)

type repository struct {
	cache cache.RedisClient
}

func NewRepository(cache cache.RedisClient) *repository {
	return &repository{
		cache: cache,
	}
}
func (r *repository) getCacheKey(uuid string) string {
	return fmt.Sprintf("%s%s", cacheKeyPrefix, uuid)
}

func (r *repository) Get(ctx context.Context, uuid string) (model.UserService, error) {
	cacheKey := r.getCacheKey(uuid)

	values, err := r.cache.HGetAll(ctx, cacheKey)
	if err != nil {
		if errors.Is(err, redigo.ErrNil) {
			return model.UserService{}, model.ErrUserNotFound
		}
		return model.UserService{}, err
	}

	if len(values) == 0 {
		return model.UserService{}, model.ErrUserNotFound
	}

	var user cacheModel.CacheUser
	err = redigo.ScanStruct(values, &user)
	if err != nil {
		return model.UserService{}, err
	}

	return converter.RedisViewerToUser(user), nil
}

func (r *repository) Set(ctx context.Context, uuid string, user model.UserService, ttl time.Duration) error {
	cacheKey := r.getCacheKey(uuid)

	redisView := converter.UserToRedisViewer(user)

	err := r.cache.HashSet(ctx, cacheKey, redisView)
	if err != nil {
		return err
	}
	return r.cache.Expire(ctx, cacheKey, ttl)
}

func (r *repository) Del(ctx context.Context, uuid string) error {
	cacheKey := r.getCacheKey(uuid)

	return r.cache.Del(ctx, cacheKey)
}
