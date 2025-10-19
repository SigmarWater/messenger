package app

import (
	"context"
	apiChat "github.com/SigmarWater/messenger/chat/internal/api/chat"
	"github.com/SigmarWater/messenger/chat/internal/closer"
	"github.com/SigmarWater/messenger/chat/internal/config"
	"github.com/SigmarWater/messenger/chat/internal/config/env"
	"github.com/SigmarWater/messenger/chat/internal/repository"
	"github.com/SigmarWater/messenger/chat/internal/repository/chats"
	"github.com/SigmarWater/messenger/chat/internal/repository/messages"
	"github.com/SigmarWater/messenger/chat/internal/service"
	serviceChat "github.com/SigmarWater/messenger/chat/internal/service/chat"
	serviceMessage "github.com/SigmarWater/messenger/chat/internal/service/message"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type serviceProvider struct {
	grpcConfig      config.GRPCConfig
	pgConfig        config.PGConfig
	pool            *pgxpool.Pool
	repoChats       repository.ChatRepository
	repoMessages    repository.MessageRepository
	srvChat         service.ChatService
	srvMessages     service.MessageService
	implChatService *apiChat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GrpcConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		grpcConfig, err := env.NewConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v\n", err)
		}

		s.grpcConfig = grpcConfig
	}
	return s.grpcConfig
}

func (s *serviceProvider) PgConfig() config.PGConfig {
	if s.pgConfig == nil {
		pgConfig, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v\n", err)
		}
		s.pgConfig = pgConfig
	}
	return s.pgConfig
}

func (s *serviceProvider) Pool(ctx context.Context) *pgxpool.Pool {
	if s.pool == nil {
		pool, err := pgxpool.Connect(ctx, s.PgConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to db: %v\n", err)
		}

		if err := pool.Ping(ctx); err != nil {
			log.Fatalf("failed to ping: %v\n", err)
		}

		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pool = pool
	}

	return s.pool
}

func (s *serviceProvider) RepChats(ctx context.Context) repository.ChatRepository {
	if s.repoChats == nil {
		s.repoChats = chats.NewPostgresChatRepository(s.Pool(ctx))
	}
	return s.repoChats
}

func (s *serviceProvider) RepMessages(ctx context.Context) repository.MessageRepository {
	if s.repoMessages == nil {
		s.repoMessages = messages.NewPostgresMessageRepository(s.Pool(ctx))
	}

	return s.repoMessages
}

func (s *serviceProvider) SrvChat(ctx context.Context) service.ChatService {
	if s.srvChat == nil {
		s.srvChat = serviceChat.NewServChat(s.RepChats(ctx))
	}
	return s.srvChat
}

func (s *serviceProvider) SrvMessages(ctx context.Context) service.MessageService {
	if s.srvMessages == nil {
		s.srvMessages = serviceMessage.NewServMessage(s.RepMessages(ctx))
	}
	return s.srvMessages
}

func (s *serviceProvider) Impl(ctx context.Context) *apiChat.Implementation {
	if s.implChatService == nil {
		s.implChatService = apiChat.NewImplementation(s.SrvChat(ctx), s.SrvMessages(ctx))
	}

	return s.implChatService
}
