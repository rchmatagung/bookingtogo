package usecase

import (
	"bookingtogo/config"
	customerModels "bookingtogo/internal/core/customer/models"
	repo "bookingtogo/internal/wrapper/repository"
	"bookingtogo/pkg/infra/db"
	"context"

	"github.com/sirupsen/logrus"
)

type Usecase interface {
	InsertCustomer(ctx context.Context, customerReq customerModels.CustomerRequest) (*customerModels.CustomerInsertResponse, error)
	GetAllCustomer(ctx context.Context) (*[]customerModels.CustomerList, error)
	GetCustomerById(ctx context.Context, customerId int64) (*customerModels.Customer, error)
}

type CustomerUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewCustomerUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) CustomerUsecase {
	return CustomerUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (c CustomerUsecase) InsertCustomer(ctx context.Context, customerReq customerModels.CustomerRequest) (*customerModels.CustomerInsertResponse, error) {

	response, err := c.Repo.Core.Customer.InsertCustomer(ctx, customerReq)
	if err != nil {
		return nil, err
	}

	return response, err

}

func (c CustomerUsecase) GetAllCustomer(ctx context.Context) (*[]customerModels.CustomerList, error) {
	var response []customerModels.CustomerList

	res, err := c.Repo.Core.Customer.GetAllCustomer(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range *res {
		response = append(response, customerModels.CustomerList{
			CustomerId:    v.CustomerId,
			CustomerName:  v.CustomerName,
			CustomerDOB:   v.CustomerDOB,
			CustomerPhone: v.CustomerPhone,
			CustomerEmail: v.CustomerEmail,
		})
	}

	return &response, err

}

func (c CustomerUsecase) GetCustomerById(ctx context.Context, customerId int64) (*customerModels.Customer, error) {

	response, err := c.Repo.Core.Customer.GetCustomerById(ctx, customerId)
	if err != nil {
		return nil, err
	}

	return response, err

}
