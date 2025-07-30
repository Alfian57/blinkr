//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Alfian57/belajar-golang/internal/handler"
	"github.com/Alfian57/belajar-golang/internal/repository"
	"github.com/Alfian57/belajar-golang/internal/service"
	"github.com/google/wire"
)

func InitializeAuthHandler() *handler.AuthHandler {
	wire.Build(handler.NewAuthHandler, service.NewAuthService, repository.NewUserRepository, repository.NewRefreshTokenRepository)
	return &handler.AuthHandler{}
}

func InitializeUserHandler() *handler.UserHandler {
	wire.Build(handler.NewUserHandler, service.NewUserService, repository.NewUserRepository)
	return &handler.UserHandler{}
}

func InitializeUrlHandler() *handler.UrlHandler {
	wire.Build(handler.NewUrlHandler, service.NewUrlService, repository.NewUrlRepository, repository.NewUserRepository)
	return &handler.UrlHandler{}
}

func InitializeUrlVisitorHandler() *handler.UrlVisitorHandler {
	wire.Build(handler.NewUrlVisitorHandler, service.NewUrlVisitorService, repository.NewUrlVisitorRepository, repository.NewUrlRepository)
	return &handler.UrlVisitorHandler{}
}

func InitializeBannedDomainHandler() *handler.BannedDomainHandler {
	wire.Build(handler.NewBannedDomainHandler, service.NewBannedDomainService, repository.NewBannedDomainRepository)
	return &handler.BannedDomainHandler{}
}

func InitializeUserService() *service.UserService {
	wire.Build(service.NewUserService, repository.NewUserRepository)
	return &service.UserService{}
}
