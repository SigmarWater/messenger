package users

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/SigmarWater/messenger/auth/internal/model"
	"github.com/SigmarWater/messenger/auth/internal/repository/converter"
	modelRepo "github.com/SigmarWater/messenger/auth/internal/repository/users/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresUserRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresUserRepository(pool *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{pool: pool}
}

func (r *PostgresUserRepository) InsertUser(ctx context.Context, user *model.UserService) (int64, error) {
	builderInsert := sq.Insert("users").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "password", "role", "create_at").
		Values(user.Name, user.Email, user.EnterPassword, user.Role, user.CreateAt).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("Ошибка при создании запроса insert: %v\n", err)
		return 0, err
	}

	var id int64
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		log.Printf("Ошибка в запросе insert к таблице users: %v\n", err)
		return 0, err
	}

	log.Printf("Добавлена запись с id в таблицу users: %d\n", id)
	return id, nil
}

func (r *PostgresUserRepository) GetUser(ctx context.Context, id int64) (*model.UserService, error) {
	builderSelect := sq.Select("name", "email", "role", "create_at", "update_at").
		PlaceholderFormat(sq.Dollar).
		From("users").
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("Ошибка при создании запроса select: %v\n", err)
		return nil, err
	}

	var userRep modelRepo.UserRepository

	if err := r.pool.QueryRow(ctx, query, args...).Scan(&userRep.Name, &userRep.Email, &userRep.Role,
		&userRep.CreateAt, &userRep.UpdateAt); err != nil {
		log.Printf("Ошибка в запросе select: %v\n", err)
		return nil, err
	}

	return converter.ToUserFromRepo(userRep), nil
}
