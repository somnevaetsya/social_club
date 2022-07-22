package usecases

import "ozon_test/app/models"

type UseCase interface {
	CreateMessage(n1 *models.Node, n2 *models.Node) error
	GetInformation() (models.Info, error)
}
