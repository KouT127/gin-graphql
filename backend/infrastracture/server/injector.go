//+build wireinject
package server

import (
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
	"github.com/KouT127/gin-sample/backend/interface/controller"
	"github.com/KouT127/gin-sample/backend/interface/gateway"
	"github.com/KouT127/gin-sample/backend/interface/presenter"
	"github.com/KouT127/gin-sample/backend/usecase/interactor"
	"github.com/google/wire"
)

var ProvideController = wire.NewSet(
	interactor.NewUserInteractor,
	controller.NewUserController,
)
var ProvideInteractor = wire.NewSet(
	ProvideController,
	gateway.NewUserRepository,
	presenter.NewUserPresenter,
)

func Inject() controller.UserController {
	wire.Build(ProvideInteractor, database.GetDB)
	return nil
}
