/*
* project ironman
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 22.02
**/

package container

import (
	"go.uber.org/dig"

	"github.com/oktopriima/ironman/application/controller/users"
)

func BuildControllerContainer(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(users.NewUserController); err != nil {
		panic(err)
	}

	return container
}
