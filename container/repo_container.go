/*
* project ironman
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 22.19
**/

package container

import (
	"go.uber.org/dig"

	"github.com/oktopriima/ironman/domain/core/services"
)

func BuildServiceContainer(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(services.NewUserServices); err != nil {
		panic(err)
	}

	return container
}
