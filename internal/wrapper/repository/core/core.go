package core

import (
	"bookingtogo/config"
	"bookingtogo/pkg/infra/db"

	customer "bookingtogo/internal/core/customer/repository"
	nationality "bookingtogo/internal/core/nationality/repository"

	"github.com/sirupsen/logrus"
)

type CoreRepository struct {
	Nationality nationality.Repository
	Customer    customer.Repository
}

func NewCoreRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreRepository {
	return CoreRepository{
		Nationality: nationality.NewNationalityRepo(dbList),
		Customer:    customer.NewCustomerRepo(dbList),
	}
}
