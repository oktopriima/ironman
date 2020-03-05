/*
* project ironman
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 21.49
**/

package users

import (
	"github.com/jinzhu/copier"

	"github.com/mkcorporate/ironman/application/request"
	"github.com/mkcorporate/ironman/domain/core/model"
)

func (u userController) CreateController(request request.UserRequest) (interface{}, error) {
	user := new(model.User)

	// copy request to user model
	if err := copier.Copy(&user, &request); err != nil {
		return nil, err
	}

	// start transaction
	tx := u.db.Begin()

	// close if any error
	defer tx.Rollback()

	m, err := u.userRepo.Create(user, tx)
	if err != nil {
		return nil, err
	}

	// commit if there's no error anymore
	tx.Commit()

	return m, nil
}
