package users

import (
	"context"
	"github.com/SigmarWater/messenger/auth/internal/model"
	"strconv"
)

func (s *serv) Get(ctx context.Context, id int64) (model.UserService, error) {
	// Сначала пытаемся получить из кеша
	uuid := strconv.FormatInt(id, 0)
	userCache, err := s.cacheRepository.Get(ctx, uuid)
	if err == nil {
		return userCache, nil
	}

	// Если нет в кеше или ошибка, идем в Postgres
	userRepo, err := s.userRepository.GetUser(ctx, id)

	if err != nil {
		return model.UserService{}, err
	}

	// Сохраняем в кеш (игнорируем ошибки кеширования)
	_ = s.cacheRepository.Set(ctx, uuid, *userRepo, s.cacheTTL)

	return *userRepo, nil
}
