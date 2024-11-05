-- name: CreateWarehouse :exec
INSERT INTO warehouses(shop_id, name, address)
VALUES (@shop_id, @name, @address);

-- name: GetWarehousesByShopID :many
SELECT w.id, w.name, w.status, w.shop_id FROM warehouses w WHERE w.shop_id = @shop_id;

-- name: AddProductToStock :exec
INSERT INTO stocks(product_id, warehouse_id, quantity)
VALUES (@product_id, @warehouse_id, @quantity);

-- name: DecreaseStock :execrows
UPDATE stocks
SET quantity = quantity - @quantity
WHERE
    product_id = @product_id AND
    warehouse_id = @warehouse_id AND
    quantity >= @quantity;

-- name: IncreaseStock :execrows
UPDATE stocks
SET quantity = quantity + @quantity
WHERE
    product_id = @product_id AND
    warehouse_id = @warehouse_id;

-- name: UpdateProductStock :exec
UPDATE stocks
SET quantity = @quantity
WHERE
    product_id = @product_id AND
    warehouse_id = @warehouse_id;

-- name: ReserveStock :exec
SELECT s.quantity
FROM stocks s
WHERE
    s.product_id = @product_id AND
    s.warehouse_id = @warehouse_id
FOR UPDATE;
