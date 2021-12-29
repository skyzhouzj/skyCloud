package main

import (
	"context"
	"fmt"
	"github.com/skyzhouzj/skyCloud/configs"
	"github.com/skyzhouzj/skyCloud/internal/router"
	"github.com/skyzhouzj/skyCloud/pkg/env"
	"github.com/skyzhouzj/skyCloud/pkg/logger"
	"github.com/skyzhouzj/skyCloud/pkg/shutdown"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	// 初始化 access logger
	accessLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.Get().SkyCloud.ProjectName, env.Active().Value())),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(configs.Get().SkyCloud.ProjectAccessLogFile),
	)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = accessLogger.Sync()
	}()

	// 初始化 HTTP 服务
	s, err := router.NewHTTPServer(accessLogger)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    configs.Get().SkyCloud.ProjectPort,
		Handler: s.Mux,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			accessLogger.Fatal("http server startup err", zap.Error(err))
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				accessLogger.Error("server shutdown err", zap.Error(err))
			}
		},

		// 关闭 db
		func() {
			if s.Db != nil {
				if err := s.Db.DbClose(); err != nil {
					accessLogger.Error("dbw close err", zap.Error(err))
				}
			}
		},

		// 关闭 cache
		func() {
			if s.Cache != nil {
				if err := s.Cache.Close(); err != nil {
					accessLogger.Error("cache close err", zap.Error(err))
				}
			}
		},
	)

}
