package sys_handler

import (
	"github.com/mojocn/base64Captcha"
	"github.com/skyzhouzj/skyCloud/configs"
	"github.com/skyzhouzj/skyCloud/internal/pkg/utils"
	"github.com/skyzhouzj/skyCloud/pkg/code"
	"github.com/skyzhouzj/skyCloud/pkg/core"
	"github.com/skyzhouzj/skyCloud/pkg/errno"
	"github.com/skyzhouzj/skyCloud/pkg/errors"
	"net/http"
)

type captchaResponse struct {
	CaptchaId string `json:"CaptchaId"` // 用户身份标识
	PicPath   string `json:"PicPath"`   // 用户身份标识
}

func (h *handler) Captcha() core.HandlerFunc {
	return func(c core.Context) {
		//字符,公式,验证码配置
		// 生成默认数字的driver
		res := new(captchaResponse)
		driver := base64Captcha.NewDriverDigit(configs.Get().Captcha.ImgHeight,
			configs.Get().Captcha.ImgWidth, configs.Get().Captcha.KeyLong, 0.7, 80)
		cp := base64Captcha.NewCaptcha(driver, utils.RedisStore.GetStore(h.cache.GetRepo()))
		if id, b64s, err := cp.Generate(); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.CaptchaError,
				code.Text(code.CaptchaError)).WithErr(errors.New(err.Error())))
		} else {
			res.CaptchaId = id
			res.PicPath = b64s
			c.Payload(res)
		}
	}
}
