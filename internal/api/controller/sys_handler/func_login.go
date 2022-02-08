package sys_handler

import (
	"github.com/skyzhouzj/skyCloud/pkg/code"
	"github.com/skyzhouzj/skyCloud/pkg/core"
	"github.com/skyzhouzj/skyCloud/pkg/errno"
	"net/http"
)

type login struct {
	Username  string `json:"userName"`
	Password  string `json:"password"`
	Captcha   string `json:"Captcha"`
	CaptchaId string `json:"CaptchaId"`
}

type userInfo struct {
	Token string `json:"Token"`
}

func (h *handler) Login() core.HandlerFunc {
	return func(c core.Context) {
		req := new(login)
		res := new(userInfo)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}
		res.Token = ""
		c.Payload(res)

	}
}
