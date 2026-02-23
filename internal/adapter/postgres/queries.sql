-- name: GetCart :one
SELECT user_id
FROM carts
WHERE user_id = $1;

-- name: GetCartItems :many
SELECT sku_id, count
FROM cart_items
WHERE user_id = $1;

-- name: EnsureCart :exec
INSERT INTO carts (user_id)
VALUES ($1)
ON CONFLICT (user_id) DO NOTHING;

-- name: DeleteCartItems :exec
DELETE
FROM cart_items
WHERE user_id = $1;

-- name: InsertCartItem :exec
INSERT INTO cart_items (user_id, sku_id, count)
VALUES ($1, $2, $3);