// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package dbgen

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Orders struct {
	ID          int64              `db:"id"`
	UserID      int64              `db:"user_id"`
	ShopID      int64              `db:"shop_id"`
	Status      string             `db:"status"`
	ProductID   int64              `db:"product_id"`
	WarehouseID int64              `db:"warehouse_id"`
	Quantity    int32              `db:"quantity"`
	TotalPrice  pgtype.Numeric     `db:"total_price"`
	CreatedAt   pgtype.Timestamptz `db:"created_at"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at"`
}

type Products struct {
	ID          int64              `db:"id"`
	Name        string             `db:"name"`
	Description string             `db:"description"`
	Price       pgtype.Numeric     `db:"price"`
	CreatedAt   pgtype.Timestamptz `db:"created_at"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at"`
}

type Shops struct {
	ID        int64              `db:"id"`
	Name      string             `db:"name"`
	Address   string             `db:"address"`
	CreatedAt pgtype.Timestamptz `db:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at"`
}

type Stocks struct {
	ID          int64              `db:"id"`
	WarehouseID int64              `db:"warehouse_id"`
	ProductID   int64              `db:"product_id"`
	Quantity    int32              `db:"quantity"`
	CreatedAt   pgtype.Timestamptz `db:"created_at"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at"`
}

type Users struct {
	ID        int64              `db:"id"`
	Email     string             `db:"email"`
	Phone     string             `db:"phone"`
	Password  string             `db:"password"`
	CreatedAt pgtype.Timestamptz `db:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at"`
}

type Warehouses struct {
	ID        int64              `db:"id"`
	ShopID    int64              `db:"shop_id"`
	Name      string             `db:"name"`
	Address   string             `db:"address"`
	Status    pgtype.Bool        `db:"status"`
	CreatedAt pgtype.Timestamptz `db:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at"`
}
