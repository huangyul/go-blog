// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/huangyul/go-blog/internal/event/article"
	"github.com/huangyul/go-blog/internal/event/history"
	"github.com/huangyul/go-blog/internal/ioc"
	"github.com/huangyul/go-blog/internal/repository"
	"github.com/huangyul/go-blog/internal/repository/cache"
	"github.com/huangyul/go-blog/internal/repository/dao"
	"github.com/huangyul/go-blog/internal/service"
	"github.com/huangyul/go-blog/internal/web"
	"github.com/huangyul/go-blog/pkg/ginx/jwt"
)

// Injectors from wire.go:

func InitApp() *App {
	cmdable := ioc.InitRedis()
	jwt := ginxjwt.NewJWT(cmdable)
	v := ioc.InitGinMiddlewares(cmdable, jwt)
	db := ioc.InitDB()
	userDAO := dao.NewUserDAOGORM(db)
	userCache := cache.NewRedisUserCache(cmdable)
	userRepository := repository.NewUserRepository(userDAO, userCache)
	userService := service.NewUserService(userRepository)
	codeCache := cache.NewRedisCodeCache(cmdable)
	codeRepository := repository.NewCodeRepository(codeCache)
	smsService := ioc.InitSMSService(cmdable)
	codeService := service.NewCodeService(codeRepository, smsService)
	userHandler := web.NewUserHandler(userService, codeService, jwt)
	articleDao := dao.NewArticleDao(db)
	articleCache := cache.NewRedisArticleCache(cmdable)
	articleRepository := repository.NewArticleRepository(articleDao, articleCache)
	client := ioc.InitSaramaClient()
	syncProducer := ioc.InitProducer(client)
	producer := article.NewSaramaSyncProducer(syncProducer)
	historyDao := dao.NewHistoryDao(db)
	historyRepository := repository.NewHistoryRepository(historyDao)
	logger := ioc.InitLogger()
	historyProducer := history.NewSaramaProducer(syncProducer, logger)
	articleService := service.NewArticleService(articleRepository, userRepository, producer, historyRepository, historyProducer, logger)
	interactiveCache := cache.NewInteractiveCache(cmdable)
	interactiveDao := dao.NewInteractiveDao(db)
	interactiveRepository := repository.NewInteractiveRepository(interactiveCache, interactiveDao)
	interactiveService := service.NewInteractiveService(interactiveRepository)
	articleHandler := web.NewArticleHandler(articleService, interactiveService)
	engine := ioc.InitServer(v, userHandler, articleHandler)
	interactiveReadConsumer := article.NewInteractiveReadConsumer(client, interactiveRepository, logger)
	consumer := history.NewConsumer(client, historyRepository, logger)
	v2 := ioc.InitConsumers(interactiveReadConsumer, consumer)
	app := &App{
		server:    engine,
		consumers: v2,
	}
	return app
}

// wire.go:

var (
	UserSet        = wire.NewSet(dao.NewUserDAOGORM, cache.NewRedisUserCache, repository.NewUserRepository, service.NewUserService, web.NewUserHandler)
	CodeSet        = wire.NewSet(repository.NewCodeRepository, cache.NewRedisCodeCache, service.NewCodeService)
	ArticleSet     = wire.NewSet(dao.NewArticleDao, cache.NewRedisArticleCache, repository.NewArticleRepository, service.NewArticleService, web.NewArticleHandler)
	InteractiveSet = wire.NewSet(dao.NewInteractiveDao, cache.NewInteractiveCache, repository.NewInteractiveRepository, service.NewInteractiveService)
	HistorySet     = wire.NewSet(dao.NewHistoryDao, repository.NewHistoryRepository)
)
