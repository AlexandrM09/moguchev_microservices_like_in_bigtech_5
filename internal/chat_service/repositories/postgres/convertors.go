package ordersrepository

import (
	"errors"
	"fmt"
	"lesson_08/orders_management_system/internal/app/models"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func convertPGError(err error) error {
	if err == nil {
		return nil
	}
	// https://github.com/jackc/pgx/wiki/Error-Handling
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		log.Println(pgErr.Message) // => syntax error at end of input
		log.Println(pgErr.Code)    // => 42601

		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			return fmt.Errorf("%s: %w", pgErr.Message, models.ErrAlreadyExists)
		default:
			return err
		}
	}
	return err
}

func newOrder(o *models.Order) (*order, error) {
	id, err := uuid.Parse(string(o.ID))
	if err != nil {
		return nil, err
	}

	return &order{
		ID:     id,
		UserID: strconv.Itoa(int(o.Buyer.ID)),
	}, nil
}
