/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 20:52
 * Copyright (c) 2019
 */

package helper

import (
	"net/http"

	"github.com/KulinaID/kulina-go-libraries/kuconfig"
	"github.com/KulinaID/kulina-go-libraries/responses"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Localization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conf := kuconfig.NewConfig()
		session := sessions.Default(ctx)

		var lang string
		v := session.Get("lang")
		if ctx.Request.Header.Get("Accept-Language") != "" {
			lang = ctx.Request.Header.Get("Accept-Language")
		} else {
			if v == nil {
				lang = conf.GetString("app.lang")
			} else {
				lang = v.(string)
			}
		}

		session.Set("lang", lang)
		err := session.Save()
		if err != nil {
			abortMission(ctx, err)
		}

		if err := InitLocales("lang", lang); err != nil {
			abortMission(ctx, err)
		}

		ctx.Next()
	}
}

func abortMission(ctx *gin.Context, err error) {
	response := new(responses.ResponseError)
	response.ErrorCode = http.StatusUnauthorized
	response.Message = err.Error()
	response.Status = http.StatusUnauthorized

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, &response)
	return
}
