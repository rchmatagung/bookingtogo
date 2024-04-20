package general

import (
	"bookingtogo/config"
	"bookingtogo/internal/wrapper/repository"
	"bookingtogo/pkg/infra/db"

	"github.com/sirupsen/logrus"
)

type GeneralUsecase struct {
}

func NewGeneralUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) GeneralUsecase {
	return GeneralUsecase{}
}
