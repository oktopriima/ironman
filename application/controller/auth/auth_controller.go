/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 22/01/2020, 19:32
 * Copyright (c) 2020
 */

package auth

import (
	"github.com/jinzhu/gorm"

	"github.com/mkcorporate/ironman/domain/core/repo"
)

// contract class
type AuthController interface {
	GoogleAuthController(request GoogleAuthRequest) (interface{}, error)
}

// define parameter
type authController struct {
	*gorm.DB
	repo.UserRepository
	repo.UserSocialRepository
}

func NewAuthController(DB *gorm.DB, userRepository repo.UserRepository, userSocialRepo repo.UserSocialRepository) AuthController {
	return &authController{DB: DB, UserRepository: userRepository, UserSocialRepository: userSocialRepo}
}
