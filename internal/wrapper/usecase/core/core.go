package core

import (
	"bookingtogo/config"
	"bookingtogo/internal/wrapper/repository"
	"bookingtogo/pkg/infra/db"

	customer "bookingtogo/internal/core/customer/usecase"
	nationality "bookingtogo/internal/core/nationality/usecase"

	"github.com/sirupsen/logrus"
)

type CoreUsecase struct {
	Nationality nationality.Usecase
	Customer    customer.Usecase
}

func NewCoreUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreUsecase {
	return CoreUsecase{
		Nationality: nationality.NewNationalityUsecase(repo, conf, dbList, log),
		Customer:    customer.NewCustomerUsecase(repo, conf, dbList, log),
	}
}
