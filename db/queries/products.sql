-- name: GetProducts :many
SELECT p.id, p.name, p.description::TEXT AS description, p.price::FLOAT AS price, s.quantity, w.shop_id, w.id AS warehouse_id
FROM products p
INNER JOIN stocks s ON p.id = s.product_id
INNER JOIN warehouses w ON s.warehouse_id = w.id AND w.status = TRUE;

-- name: GetProductsByShopID :many
SELECT p.id, p.name, p.description::TEXT AS description, p.price::FLOAT AS price, s.quantity, w.shop_id, s.warehouse_id
FROM products p
INNER JOIN stocks s ON p.id = s.product_id
INNER JOIN warehouses w ON s.warehouse_id = w.id AND w.status = TRUE
WHERE w.shop_id = @shop_id;

-- name: AddProduct :one
INSERT INTO products(name, description, price)
VALUES (@name, @description, @price)
RETURNING id;

-- name: GetProductByID :one
SELECT p.id, p.name, p.description::TEXT AS description, p.price::FLOAT AS price
FROM products p
INNER JOIN stocks s ON p.id = s.product_id
INNER JOIN warehouses w ON s.warehouse_id = w.id AND w.status = TRUE
WHERE p.id = @id;
