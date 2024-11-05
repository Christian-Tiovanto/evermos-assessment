-- name: CreateShop :exec
INSERT INTO shops(name, address) VALUES (@name, @address);
