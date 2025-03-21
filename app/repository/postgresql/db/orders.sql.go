// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: orders.sql

package dbgen

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOrder = `-- name: CreateOrder :one
Insert into orders (user_id, shop_id, product_id, warehouse_id, quantity, total_price)
Values ($1, $2, $3, $4, $5, $6)
Returning id
`

type CreateOrderParams struct {
	UserID      int64          `db:"user_id"`
	ShopID      int64          `db:"shop_id"`
	ProductID   int64          `db:"product_id"`
	WarehouseID int64          `db:"warehouse_id"`
	Quantity    int32          `db:"quantity"`
	TotalPrice  pgtype.Numeric `db:"total_price"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (int64, error) {
	row := q.db.QueryRow(ctx, createOrder,
		arg.UserID,
		arg.ShopID,
		arg.ProductID,
		arg.WarehouseID,
		arg.Quantity,
		arg.TotalPrice,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateOrderStatus = `-- name: UpdateOrderStatus :one
UPDATE orders
SET status = CASE
    WHEN $1 = 'PAID' THEN 'PAID'
    WHEN $1 = 'CANCELLED' THEN 'CANCELLED'
    ELSE status
END
WHERE id = $2 AND status = 'PENDING'
RETURNING product_id, warehouse_id, quantity, status
`

type UpdateOrderStatusParams struct {
	Status interface{} `db:"status"`
	ID     int64       `db:"id"`
}

type UpdateOrderStatusRow struct {
	ProductID   int64  `db:"product_id"`
	WarehouseID int64  `db:"warehouse_id"`
	Quantity    int32  `db:"quantity"`
	Status      string `db:"status"`
}

func (q *Queries) UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) (UpdateOrderStatusRow, error) {
	row := q.db.QueryRow(ctx, updateOrderStatus, arg.Status, arg.ID)
	var i UpdateOrderStatusRow
	err := row.Scan(
		&i.ProductID,
		&i.WarehouseID,
		&i.Quantity,
		&i.Status,
	)
	return i, err
}
