package converter

import (
	"github.com/SigmarWater/messenger/auth/internal/model"
	modelRepo "github.com/SigmarWater/messenger/auth/internal/repository/model"
	"strconv"
	"time"
)

func ToUserFromRepo(user modelRepo.UserRepository) *model.UserService {
	return &model.UserService{
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		CreateAt: user.CreateAt.Time,
		UpdateAt: user.UpdateAt.Time,
	}
}

// Конвертер из сервисной модели в redis модель
func UserToRedisViewer(user model.UserService) modelRepo.CacheUser {
	var uuid string
	uuid = strconv.FormatInt(int64(user.Id), 10)

	var createAt int64
	createAt = user.CreateAt.UnixNano()

	var updateAt int64
	updateAt = user.CreateAt.UnixNano()

	return modelRepo.CacheUser{
		UUID:     uuid,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		CreateAt: createAt,
		UpdateAt: updateAt,
	}
}

func RedisViewerToUser(user modelRepo.CacheUser) *model.UserService {
	var id int64
	id, err := strconv.ParseInt(user.UUID, 0, 64)
	if err != nil {
		return nil
	}

	var createAt time.Time
	createAt = time.Unix(0, user.CreateAt)

	var updateAt time.Time
	updateAt = time.Unix(0, user.UpdateAt)

	return &model.UserService{
		Id:       int(id),
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		CreateAt: createAt,
		UpdateAt: updateAt,
	}
}
