/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/11/2019, 19:53
 * Copyright (c) 2019
 */

package services

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/mkcorporate/ironman/domain/core/model"
	"github.com/mkcorporate/ironman/domain/core/repo"
)

type userServices struct {
	db *gorm.DB
}

func NewUserServices(db *gorm.DB) repo.UserRepository {
	return &userServices{db}
}

func (srv *userServices) Create(user *model.User, tx *gorm.DB) (*model.User, error) {
	db := tx.Create(&user)
	m := new(model.User)

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
