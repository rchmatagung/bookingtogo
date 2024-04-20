package repository

import (
	"bookingtogo/internal/core/nationality/models"
	"bookingtogo/pkg/infra/db"
	"context"
)

type Repository interface {
	InsertNationality(ctx context.Context, nationalityReq models.NationalityRequest) (*models.Nationality, error)
	GetAllNationality(ctx context.Context) (*[]models.Nationality, error)
}

type NationalityRepo struct {
	DBList *db.DatabaseList
}

func NewNationalityRepo(dbList *db.DatabaseList) NationalityRepo {
	return NationalityRepo{
		DBList: dbList,
	}
}

func (n NationalityRepo) InsertNationality(ctx context.Context, nationalityReq models.NationalityRequest) (*models.Nationality, error) {
	var response models.Nationality

	err := n.DBList.DatabaseApp.Raw(InsertNationality, nationalityReq.NationalityName, nationalityReq.NationalityCode).Scan(&response).Error

	return &response, err

}

func (n NationalityRepo) GetAllNationality(ctx context.Context) (*[]models.Nationality, error) {
	var response []models.Nationality

	err := n.DBList.DatabaseApp.Raw(SelectAllNationality).Scan(&response).Error

	return &response, err

}
