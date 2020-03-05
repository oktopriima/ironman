/*
* project ironman
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 21.45
**/

package users

import (
	"github.com/jinzhu/gorm"

	"github.com/mkcorporate/ironman/application/request"
	"github.com/mkcorporate/ironman/domain/core/repo"
)

type UserController interface {
	CreateController(request request.UserRequest) (interface{}, error)
}

type userController struct {
	db       *gorm.DB
	userRepo repo.UserRepository
}

func NewUserController(db *gorm.DB,
	userRepo repo.UserRepository) UserController {
	return &userController{db, userRepo}
}
