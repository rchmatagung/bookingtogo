package usecase

import (
	"bookingtogo/config"
	"bookingtogo/internal/wrapper/repository"
	"bookingtogo/internal/wrapper/usecase/core"
	"bookingtogo/internal/wrapper/usecase/general"
	"bookingtogo/pkg/infra/db"

	"github.com/sirupsen/logrus"
)

type Usecase struct {
	General general.GeneralUsecase
	Core    core.CoreUsecase
}

func NewUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) Usecase {
	return Usecase{
		General: general.NewGeneralUsecase(repo, conf, dbList, log),
		Core:    core.NewCoreUsecase(repo, conf, dbList, log),
	}
}
