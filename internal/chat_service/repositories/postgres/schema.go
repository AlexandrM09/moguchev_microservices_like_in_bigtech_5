package ordersrepository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const tableOrders = "public.orders"

const (
	columnID     = "order_id"
	columnUserID = "user_id"
)

type order struct {
	ID        uuid.UUID           `db:"order_id"`
	UserID    string              `db:"user_id"`
	UpdatedAt sql.Null[time.Time] `db:"updated_at"`
}

var (
	tableOrdersColumns = []string{
		columnID,
		columnUserID,
	}
)

func (o *order) mapFields() map[string]any {
	return map[string]any{
		columnID:     o.ID,
		columnUserID: o.UserID,
	}
}

func (o *order) Values(columns ...string) []any {
	mapFields := o.mapFields()
	values := make([]any, 0, len(columns))
	for i := range columns {
		if v, ok := mapFields[columns[i]]; ok {
			values = append(values, v)
		} else {
			values = append(values, nil)
		}
	}
	return values
}
