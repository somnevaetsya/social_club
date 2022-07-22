package usecases_impl

import (
	"social_club/app/models"
	"social_club/app/repositories"
	"social_club/app/usecases"
	"social_club/pkg/errors"
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
	isEmpty, err := useCase.rep.IsEmpty()
	if err != nil {
		return models.Info{}, err
	} else if isEmpty == true {
		return models.Info{IsEmpty: true}, err
	}
	info, err := useCase.rep.GetGraph()
	if err != nil {
		return models.Info{}, err
	}
	return info, nil
}
