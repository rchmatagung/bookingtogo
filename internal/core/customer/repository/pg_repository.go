package repository

import (
	"bookingtogo/internal/core/customer/models"
	"bookingtogo/pkg/infra/db"
	"context"
)

type Repository interface {
	InsertCustomer(ctx context.Context, customerReq models.CustomerRequest) (*models.CustomerInsertResponse, error)
	GetAllCustomer(ctx context.Context) (*[]models.CustomerList, error)
	GetCustomerById(ctx context.Context, customerId int64) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, customerReq models.UpdateCustomerRequest) (*models.CustomerInsertResponse, error)
}

type CustomerRepo struct {
	DBList *db.DatabaseList
}

func NewCustomerRepo(dbList *db.DatabaseList) CustomerRepo {
	return CustomerRepo{
		DBList: dbList,
	}
}

func (c CustomerRepo) InsertCustomer(ctx context.Context, customerReq models.CustomerRequest) (*models.CustomerInsertResponse, error) {

	var response models.CustomerInsertResponse
	var familyResponse models.FamilyList

	err := c.DBList.DatabaseApp.Raw(InsertCustomer, customerReq.NationalityId, customerReq.CustomerName, customerReq.CustomerDOB, customerReq.CustomerPhone, customerReq.CustomerEmail).Scan(&response).Error

	for _, v := range customerReq.FamilyList {
		err = c.DBList.DatabaseApp.Raw(InsertFamilyList, response.CustomerId, v.FamilyListRelation, v.FamilyListName, v.FamilyListDOB).Scan(&familyResponse).Error
		if err != nil {
			return nil, err
		}

		response.FamilyList = append(response.FamilyList, familyResponse)
	}

	return &response, err

}

func (c CustomerRepo) GetAllCustomer(ctx context.Context) (*[]models.CustomerList, error) {

	var response []models.CustomerList

	err := c.DBList.DatabaseApp.Raw(SelectAllCustomer).Scan(&response).Error

	return &response, err

}

func (c CustomerRepo) GetCustomerById(ctx context.Context, customerId int64) (*models.Customer, error) {

	var customerRes models.Customer
	var familyList []models.FamilyList

	err := c.DBList.DatabaseApp.Raw(SelectCustomerById, customerId).Scan(&customerRes).Error
	if err != nil {
		return nil, err
	}

	err = c.DBList.DatabaseApp.Raw(SelectFamilyListByCustomerId, customerId).Scan(&familyList).Error
	if err != nil {
		return nil, err
	}

	customerRes.FamilyList = familyList

	return &customerRes, err

}

func (c CustomerRepo) UpdateCustomer(ctx context.Context, customerReq models.UpdateCustomerRequest) (*models.CustomerInsertResponse, error) {

	var response models.CustomerInsertResponse
	var familyResponse models.FamilyList
	var err error

	err = c.DBList.DatabaseApp.Raw(UpdateCustomer, customerReq.NationalityId, customerReq.CustomerName, customerReq.CustomerDOB, customerReq.CustomerPhone, customerReq.CustomerEmail, customerReq.CustomerId).Scan(&response).Error
	if err != nil {
		return nil, err
	}

	for _, v := range customerReq.FamilyList {
		if v.FamilyListId == 0 {
			err = c.DBList.DatabaseApp.Raw(InsertFamilyList, response.CustomerId, v.FamilyListRelation, v.FamilyListName, v.FamilyListDOB).Scan(&familyResponse).Error
			if err != nil {
				return nil, err
			}
		} else {
			err = c.DBList.DatabaseApp.Raw(UpdateFamilyList, v.FamilyListRelation, v.FamilyListName, v.FamilyListDOB, v.FamilyListId).Scan(&familyResponse).Error
			if err != nil {
				return nil, err
			}
		}

		if v.IsDelete {
			err = c.DBList.DatabaseApp.Exec(DeleteFamilyList, v.FamilyListId).Error
			if err != nil {
				return nil, err
			}
		}

		response.FamilyList = append(response.FamilyList, familyResponse)
	}

	return &response, err

}
