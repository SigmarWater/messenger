package users

import (
	"context"
	"github.com/SigmarWater/messenger/auth/internal/model"
	"strconv"
)

func (s *serv) Create(ctx context.Context, user *model.UserService) (int64, error) {
	// Создаем в Postgres
	id, err := s.userRepository.InsertUser(ctx, user)

	if err != nil {
		return 0, err
	}

	uuid := strconv.FormatInt(id, 0)
	// Получаем созданную запись для кеширования
	userRepo, err := s.userRepository.GetUser(ctx, id)
	if err == nil {
		// Кешируем созданную запись (игнорируем ошибки кеширования)
		_ = s.cacheRepository.Set(ctx, uuid, *userRepo, s.cacheTTL)
	}

	return id, nil
}
