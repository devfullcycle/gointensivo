package database

import "github.com/devfullcycle/gointensivo/internal/order/entity"

type OrderRepositoryInterface interface {
	Save(order *entity.Order) error
	GetTotal() (int, error)
}
