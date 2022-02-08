package router

import "github.com/skyzhouzj/skyCloud/internal/api/controller/sys_handler"

func setApiRouter(r *resource) {

	sysHandler := sys_handler.New(r.logger, r.db, r.cache)
	login := r.mux.Group("/p/c")
	{
		login.POST("/Login", sysHandler.Login())
		login.GET("/Captcha", sysHandler.Captcha())
	}

}
