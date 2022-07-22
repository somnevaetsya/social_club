package repositories

import "ozon_test/app/models"

type Repository interface {
	Add(n1 *models.Node, n2 *models.Node) error
	GetInfo() (min, max uint, avg float32, err error)
	GetGraph() (models.Info, error)
}
