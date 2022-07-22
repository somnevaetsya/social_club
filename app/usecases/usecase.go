package usecases

import "social_club/app/models"

type UseCase interface {
	CreateMessage(n1 *models.Node, n2 *models.Node) error
	GetInformation() (models.Info, error)
}
