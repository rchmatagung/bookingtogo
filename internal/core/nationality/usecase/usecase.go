package usecase

import (
	"bookingtogo/config"
	"bookingtogo/internal/core/nationality/models"
	repo "bookingtogo/internal/wrapper/repository"
	"bookingtogo/pkg/infra/db"
	"context"

	"github.com/sirupsen/logrus"
)

type Usecase interface {
	InsertNationality(ctx context.Context, nationalityReq models.NationalityRequest) (*models.Nationality, error)
	GetAllNationality(ctx context.Context) (*[]models.Nationality, error)
}

type NationalityUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewNationalityUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) NationalityUsecase {
	return NationalityUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (n NationalityUsecase) InsertNationality(ctx context.Context, nationalityReq models.NationalityRequest) (*models.Nationality, error) {

	response, err := n.Repo.Core.Nationality.InsertNationality(ctx, nationalityReq)
	if err != nil {
		return nil, err
	}

	return response, err

}

func (n NationalityUsecase) GetAllNationality(ctx context.Context) (*[]models.Nationality, error) {

	response, err := n.Repo.Core.Nationality.GetAllNationality(ctx)
	if err != nil {
		return nil, err
	}

	return response, err
}
