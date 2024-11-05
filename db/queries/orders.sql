-- name: CreateOrder :one
Insert into orders (user_id, shop_id, product_id, warehouse_id, quantity, total_price)
Values (@user_id, @shop_id, @product_id, @warehouse_id, @quantity, @total_price)
Returning id;

-- name: UpdateOrderStatus :one
UPDATE orders
SET status = CASE
    WHEN @status = 'PAID' THEN 'PAID'
    WHEN @status = 'CANCELLED' THEN 'CANCELLED'
    ELSE status
END
WHERE id = @id AND status = 'PENDING'
RETURNING product_id, warehouse_id, quantity, status;
