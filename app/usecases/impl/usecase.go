package usecases_impl

import (
	"ozon_test/app/models"
	"ozon_test/app/repositories"
	"ozon_test/app/usecases"
	"ozon_test/pkg/errors"
)

type UseCaseImpl struct {
	rep repositories.Repository
}

func MakeUseCase(rep_ repositories.Repository) usecases.UseCase {
	return &UseCaseImpl{rep: rep_}
}

func (useCase *UseCaseImpl) CreateMessage(n1 *models.Node, n2 *models.Node) error {
	err := useCase.rep.Add(n1, n2)
	if err != nil {
		return customErrors.ErrBadInputData
	}
	return nil
}

func (useCase *UseCaseImpl) GetInformation() (models.Info, error) {
	info, err := useCase.rep.GetGraph()
	if err != nil {
		return models.Info{}, err
	}
	return info, nil
}
