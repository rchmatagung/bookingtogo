package core

import (
	"bookingtogo/config"
	"bookingtogo/internal/wrapper/usecase"

	customer "bookingtogo/internal/core/customer/delivery"
	nationality "bookingtogo/internal/core/nationality/delivery"

	"github.com/sirupsen/logrus"
)

type CoreHandler struct {
	Nationality nationality.NationalityHandler
	Customer    customer.CustomerHandler
}

func NewCoreHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) CoreHandler {
	return CoreHandler{
		Nationality: nationality.NewNationalityHandler(uc, conf, log),
		Customer:    customer.NewCustomerHandler(uc, conf, log),
	}
}
