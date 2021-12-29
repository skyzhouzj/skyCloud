package router

import (
	"github.com/skyzhouzj/skyCloud/pkg/cache"
	"github.com/skyzhouzj/skyCloud/pkg/core"
	"github.com/skyzhouzj/skyCloud/pkg/db"
	"github.com/skyzhouzj/skyCloud/pkg/errors"
	"github.com/skyzhouzj/skyCloud/pkg/middleware"
	"go.uber.org/zap"
)

type resource struct {
	mux     core.Mux
	logger  *zap.Logger
	db      db.Repo
	cache   cache.Repo
	middles middleware.Middleware
}

type Server struct {
	Mux   core.Mux
	Db    db.Repo
	Cache cache.Repo
}

func NewHTTPServer(logger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	r := new(resource)
	r.logger = logger

	// 初始化 DB
	dbRepo, err := db.New()
	if err != nil {
		logger.Fatal("new db err", zap.Error(err))
	}
	r.db = dbRepo

	// 初始化 Cache
	cacheRepo, err := cache.New()
	if err != nil {
		logger.Fatal("new cache err", zap.Error(err))
	}
	r.cache = cacheRepo

	mux, err := core.New(logger,
		core.WithEnableCors(),
		core.WithEnableRate(),
	)

	if err != nil {
		panic(err)
	}

	r.mux = mux
	r.middles = middleware.New(logger, r.cache, r.db)

	// 设置 API 路由
	setApiRouter(r)

	s := new(Server)
	s.Mux = mux
	s.Db = r.db
	s.Cache = r.cache

	return s, nil
}
