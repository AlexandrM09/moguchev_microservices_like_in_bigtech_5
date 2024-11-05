package ordersrepository

import (
	"lesson_08/orders_management_system/internal/app/usecases/orders"
	"lesson_08/orders_management_system/pkg/postgres/transaction_manager"
)

type Repository struct {
	txManager transaction_manager.TransactionManagerAPI
}

var (
	_ orders.OrdersRepository = (*Repository)(nil)
)

func NewRepository(txManager transaction_manager.TransactionManagerAPI) *Repository {
	return &Repository{
		txManager: txManager,
	}
}
