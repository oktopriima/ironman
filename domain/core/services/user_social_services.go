/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 22/01/2020, 20:22
 * Copyright (c) 2020
 */

package services

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/mkcorporate/ironman/domain/core/model"
	"github.com/mkcorporate/ironman/domain/core/repo"
)

type userSocialServices struct {
	db *gorm.DB
}

func (u userSocialServices) Create(social *model.UserSocial, tx *gorm.DB) (*model.UserSocial, error) {
	db := tx.Create(&social)
	m := new(model.UserSocial)

	if err := db.Error; err != nil {
		return nil, err
	}

	byteData, err := json.Marshal(db.Value)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(byteData, &m); err != nil {
		return nil, err
	}

	return m, nil
}

func (u userSocialServices) Update(social *model.UserSocial, tx *gorm.DB) error {
	panic("implement me")
}

func (u userSocialServices) FindOneBy(criteria map[string]interface{}) (*model.UserSocial, error) {
	m := new(model.UserSocial)
	if err := u.db.Where(criteria).Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (u userSocialServices) FindPagedBy(criteria map[string]interface{}, page, size int) ([]*model.UserSocial, error) {
	panic("implement me")
}

func NewUserSocialServices(db *gorm.DB) repo.UserSocialRepository {
	return &userSocialServices{db: db}
}
