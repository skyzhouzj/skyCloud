package middleware

import (
	"github.com/skyzhouzj/skyCloud/configs"
	"github.com/skyzhouzj/skyCloud/pkg/code"
	"github.com/skyzhouzj/skyCloud/pkg/core"
	"github.com/skyzhouzj/skyCloud/pkg/errno"
	"github.com/skyzhouzj/skyCloud/pkg/errors"
	"github.com/skyzhouzj/skyCloud/pkg/signature"
	"github.com/skyzhouzj/skyCloud/pkg/urltable"
	"net/http"
	"strings"
	"time"
)

const (
	ttl       = time.Minute * 10 // 签名超时时间 10 分钟
	minLength = 2                // split space 最小长度
	notUsed   = -1               // -1 表示被禁用
)

var whiteListPath = map[string]bool{
	"/login": true,
}

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

		key := authorizationSplit[0]

		data, err := m.authorizedService.DetailByKey(c, key)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithErr(err),
			)
			return
		}

		if data.IsUsed == notUsed {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithErr(errors.New(key + " 已被禁止调用")),
			)
			return
		}

		if len(data.Apis) < 1 {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithErr(errors.New(key + " 未进行接口授权")),
			)
			return
		}

		if !whiteListPath[c.Path()] {
			// 验证 c.Method() + c.Path() 是否授权
			table := urltable.NewTable()
			for _, v := range data.Apis {
				_ = table.Append(v.Method + v.Api)
			}

			if pattern, _ := table.Mapping(c.Method() + c.Path()); pattern == "" {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.AuthorizationError,
					code.Text(code.AuthorizationError)).WithErr(errors.New(c.Method() + c.Path() + " 未进行接口授权")),
				)
				return
			}
		}

		ok, err := signature.New(key, data.Secret, ttl).Verify(authorization, date, c.Path(), c.Method(), c.RequestInputParams())
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithErr(err),
			)
			return
		}

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
