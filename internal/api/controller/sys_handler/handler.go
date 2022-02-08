package sys_handler

import (
	"github.com/skyzhouzj/skyCloud/configs"
	"github.com/skyzhouzj/skyCloud/pkg/cache"
	"github.com/skyzhouzj/skyCloud/pkg/core"
	"github.com/skyzhouzj/skyCloud/pkg/db"
	"github.com/skyzhouzj/skyCloud/pkg/hash"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Login 登录
	// @Tags API.sys.login
	// @Router /p/c/Login [post]
	Login() core.HandlerFunc

	// Logout 登出
	// @Tags API.sys.loginout
	// @Router /p/cs/Logout [post]
	Logout() core.HandlerFunc

	// Captcha 验证码
	// @Tags API.sys.Captcha
	// @Router /p/c/Captcha [post]
	Captcha() core.HandlerFunc
}

type handler struct {
	logger  *zap.Logger
	cache   cache.Repo
	hashids hash.Hash
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:  logger,
		cache:   cache,
		hashids: hash.New(configs.Get().SkyCloud.HashIds.Secret, configs.Get().SkyCloud.HashIds.Length),
	}
}

func (h *handler) i() {}
