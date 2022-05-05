//go:build InitializeServer
// +build InitializeServer

package cenco_pim

import (
	"cenco-pim/adapter/dependecy"
	"cenco-pim/app/handler"
	"cenco-pim/app/repository"
	"cenco-pim/app/router"
	"cenco-pim/app/service"
	"cenco-pim/util/logger"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeServer(dbConn *gorm.DB, logger *logger.Logger) dependecy.InitDependencyImpl {
	wire.Build(
		repository.NewFamilyRepository,

		service.NewFamilyServiceImpl,

		handler.NewFamilyHandlerImpl,

		router.NewFamily,

		dependecy.NewInitDependencyImpl,
	)
	return dependecy.InitDependencyImpl{}
}
