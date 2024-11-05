package ordersrepository

import (
	"context"

	orders_models "lesson_08/orders_management_system/internal/app/usecases/orders/models"
	"lesson_08/orders_management_system/pkg/pagination"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) FindOrders(ctx context.Context, filter *orders_models.OrdersSearchFilter, p *pagination.Pagination) (*orders_models.ListOrdersResponse, error) {
	qb := squirrel.Select(tableOrdersColumns...).
		From(tableOrders).
		PlaceholderFormat(squirrel.Dollar)

	qb = applyOrdersSearchFilter(qb, filter)
	qb = applyPagination(qb, p)

	var orders []order
	err := r.txManager.GetQueryEngine(ctx).Selectx(ctx, &orders, qb)
	if err != nil {
		return nil, err
	}

	// конвернтация

	return nil, nil
}

func applyOrdersSearchFilter(qb squirrel.SelectBuilder, filter *orders_models.OrdersSearchFilter) squirrel.SelectBuilder {
	if filter == nil {
		return qb
	}

	if len(filter.OrderIDs) > 0 {
		qb = qb.Where(squirrel.Eq{columnID: filter.OrderIDs})
	}

	if len(filter.UserIDs) > 0 {
		qb = qb.Where(squirrel.Eq{columnUserID: filter.UserIDs})
	}

	return qb
}

func applyPagination(qb squirrel.SelectBuilder, p *pagination.Pagination) squirrel.SelectBuilder {
	if p == nil {
		return qb
	}

	return qb.Limit(p.PageSize).Offset(p.PageSize * (p.PageNumber - 1))
}
