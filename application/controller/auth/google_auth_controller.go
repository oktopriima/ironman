/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 22/01/2020, 20:06
 * Copyright (c) 2020
 */

package auth

import (
	"github.com/gosimple/slug"

	"github.com/mkcorporate/ironman/domain/core/model"
	"github.com/mkcorporate/ironman/domain/helper"
)

type GoogleAuthRequest interface {
	GetToken() string
}

func (uc *authController) GoogleAuthController(request GoogleAuthRequest) (interface{}, error) {

	gResp, err := helper.GetGoogleData(request.GetToken())
	if err != nil {
		return nil, err
	}

	social, err := uc.UserSocialRepository.FindOneBy(map[string]interface{}{
		"social_id": gResp.ID,
	})

	if err == nil && social != nil {
		//	user does not exist
		user := new(model.User)
		user.Name = gResp.DisplayName
		user.Username = slug.MakeLang(gResp.DisplayName, "en")
		user.Email = gResp.Emails[0].Value
		user.Phone = ""

	}

	return gResp, nil
}
