//+build wireinject
package server

import (
	"github.com/KouT127/gin-sample/backend/application/interactor"
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
	"github.com/KouT127/gin-sample/backend/infrastracture/datastore"
	"github.com/KouT127/gin-sample/backend/interface/controller"
	"github.com/KouT127/gin-sample/backend/interface/presenter"
)

func InjectUser() *controller.UserController {
	db := database.NewDB()
	ur := datastore.NewUserRepository(db)
	up := presenter.NewUserPresenter()
	ui := interactor.NewUserInteractor(ur, up)
	uc := controller.NewUserController(ui)
	return uc
}
