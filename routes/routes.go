/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 21:33
 * Copyright (c) 2019
 */

package routes

import (
	"github.com/KulinaID/kulina-go-libraries/kumiddleware"
	"github.com/gin-gonic/gin"

	"github.com/oktopriima/ironman/application/httphandler/users"
	"github.com/oktopriima/ironman/domain/config"
)

func InvokeRoute(
	engine *gin.Engine,
	user users.UserHandler,
) {
	conf := config.NewConfig()
	route := engine.Group("api/" + conf.GetString("app.version.tag") + conf.GetString("app.version.value"))

	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	route.Use(gin.ErrorLogger())

	route.OPTIONS("/*path", kumiddleware.CORSMiddleware())

	// user route
	{
		userRoute := route.Group("users")
		userRoute.POST("", user.CreateHandler)
	}

}
