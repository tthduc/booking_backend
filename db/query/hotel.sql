-- name: CreateHotel :one
INSERT INTO hotel (name,
                   address,
                   location) VALUES (
   $1, $2, $3
) RETURNING *;

-- name: GetHotel :one
SELECT * FROM hotel
WHERE id = $1 LIMIT 1;

-- name: ListHotels :many
SELECT * FROM hotel
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: UpdateHotel :one
UPDATE hotel
SET name = $2, address = $3, location = $4
WHERE id = $1
RETURNING *;

-- name: DeleteHotel :exec
DELETE FROM hotel
WHERE id = $1;