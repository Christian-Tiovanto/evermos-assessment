-- name: CreateUser :exec
INSERT INTO "users"("email", "phone", "password")
VALUES (@email, @phone, @password);

-- name: GetUserByEmailOrPhone :one
SELECT * FROM "users" WHERE "email" = @email OR "phone" = @phone;

-- name: GetUsers :many
SELECT id, email, phone FROM "users" ORDER BY id;

-- name: UpdateUserById :exec
UPDATE "users"
SET 
    "email" = COALESCE(@email, "email"),
    "phone" = COALESCE(@phone, "phone"),
    "password" = COALESCE(@password, "password")
WHERE "id" = @id;
