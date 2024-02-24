-- name: CreateRoomInventory :one
INSERT INTO room_inventory (hotel_id,
                   room_type_id,
                   date,
                   total_inventory,
                   total_reserved) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: UpdateRoomInventory :one
UPDATE room_inventory
SET
    date = $3,
    total_inventory = total_inventory - 1,
    total_reserved = total_reserved + 1
WHERE hotel_id = $1 and room_type_id = $2
RETURNING *;

-- name: ListRoomInventory :many
SELECT * FROM room_inventory
WHERE hotel_id = $3
ORDER BY hotel_id
LIMIT $1
OFFSET $2;

-- name: GetRoomInventory :one
SELECT * FROM room_inventory
WHERE hotel_id = $1 AND room_type_id = $2;