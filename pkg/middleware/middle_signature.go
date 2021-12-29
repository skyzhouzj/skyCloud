package middleware

import (
	"github.com/skyzhouzj/skyCloud/configs"
	"github.com/skyzhouzj/skyCloud/pkg/code"
	"github.com/skyzhouzj/skyCloud/pkg/core"
	"github.com/skyzhouzj/skyCloud/pkg/errno"
	"github.com/skyzhouzj/skyCloud/pkg/errors"
	"net/http"
	"strings"
	"time"
)

const (
	ttl       = time.Minute * 10 // 签名超时时间 10 分钟
	minLength = 2                // split space 最小长度
	notUsed   = -1               // -1 表示被禁用

)

func (m *middleware) Signature() core.HandlerFunc {
	return func(c core.Context) {
		// 签名信息
		authorization := c.GetHeader(configs.Get().SkyCloud.HeaderSignToken)
		if authorization == "" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithErr(errors.New("Header 中缺少 XhFramwork 参数")),
			)
			return
		}

		// 时间信息
		date := c.GetHeader(configs.Get().SkyCloud.HeaderSignTokenDate)
		if date == "" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithErr(errors.New("Header 中缺少 Date 参数")),
			)
			return
		}

		// 通过签名信息获取 key
		authorizationSplit := strings.Split(authorization, " ")
		if len(authorizationSplit) < minLength {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithErr(errors.New("Header 中 XhFramwork 格式错误")),
			)
			return
		}

		ok := (len(authorizationSplit[0]) > 0)

		if !ok {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithErr(errors.New("Header 中 XhFramwork 信息错误")),
			)
			return
		}
	}
}
