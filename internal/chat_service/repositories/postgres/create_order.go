package ordersrepository

import (
	"context"
	"fmt"

	"lesson_08/orders_management_system/internal/app/models"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) CreateOrder(ctx context.Context, order *models.Order) error {
	row, err := newOrder(order)
	if err != nil {
		return err
	}

	qb := squirrel.Insert(tableOrders).
		Columns(tableOrdersColumns...).
		Values(row.Values(tableOrdersColumns...)...).
		PlaceholderFormat(squirrel.Dollar)

	_, err = r.txManager.GetQueryEngine(ctx).Execx(ctx, qb)
	if err != nil {
		return fmt.Errorf("postgres: %w", convertPGError(err))
	}

	return nil
}
