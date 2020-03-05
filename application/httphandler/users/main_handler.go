/*
* project ironman
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 22.03
**/

package users

import (
	"github.com/gin-gonic/gin"

	"github.com/oktopriima/ironman/application/controller/users"
)

type UserHandler interface {
	CreateHandler(ctx *gin.Context)
}

type userHandler struct {
	user users.UserController
}

func NewUserHandler(user users.UserController) UserHandler {
	return &userHandler{user}
}
