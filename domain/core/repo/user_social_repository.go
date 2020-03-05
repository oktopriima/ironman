/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 22/01/2020, 20:17
 * Copyright (c) 2020
 */

package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/mkcorporate/ironman/domain/core/model"
)

type UserSocialRepository interface {
	Create(social *model.UserSocial, tx *gorm.DB) (*model.UserSocial, error)
	Update(social *model.UserSocial, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.UserSocial, error)
	FindPagedBy(criteria map[string]interface{}, page, size int) ([]*model.UserSocial, error)
}
