package bootstrap

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joqd/slovo/internal/adapter/config"
	"github.com/joqd/slovo/internal/adapter/controller/http"
	"github.com/joqd/slovo/internal/adapter/repository/cache"
	"github.com/joqd/slovo/internal/adapter/repository/persistent"
	"github.com/joqd/slovo/internal/core/usecase"
	"github.com/joqd/slovo/pkg/httpserver"
	"github.com/joqd/slovo/pkg/logger"
	"github.com/joqd/slovo/pkg/mongodb"
	"github.com/joqd/slovo/pkg/redisdb"
)

func Run(conf *config.Config) {
	// Initialize logger
	xlog := logger.New(conf.Log.Level)

	// Initialize MongoDB
	mongoDB, err := mongodb.New(conf.Mongo.URI, conf.Mongo.DbName)
	if err != nil {
		xlog.Fatal("Failed to connect to MongoDB: %v", err)
	}

	// Initialize Redis
	redisDB, err := redisdb.New(
		conf.Redis.URI,
		redisdb.WithTTL(time.Duration(conf.Redis.TTL)*time.Hour),
	)
	if err != nil {
		xlog.Fatal("Failed to connect to Redis: %v", err)
	}

	// Initialize Repositories
	wordPersistentRepo := persistent.NewWordRespository(*mongoDB, xlog)
	wordCacheRepo := cache.NewWordCache(*redisDB, xlog)

	// Initialize Usecases
	wordUsecase := usecase.NewWordUsecase(wordPersistentRepo, wordCacheRepo, xlog)

	// Initialize Server
	server := httpserver.New()

	// Initialize HTTP
	httpOptions := http.Options{
		Engine:      server.Engine,
		Conf:        conf,
		Log:         xlog,
		WordUsecase: wordUsecase,
	}
	http.RegisterRoutes(httpOptions)

	// Start HTTP server
	server.Start()

	// Wait for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := server.Shutdown(); err != nil {
		xlog.Fatal("The server was forcibly shut down: %v", err)
	}
}
