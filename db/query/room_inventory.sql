-- name: CreateRoomInventory :one
INSERT INTO room_inventory (hotel_id,
                   room_id,
                   room_type_id,
                   date,
                   total_inventory,
                   total_reserved) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: UpdateRoomInventory :one
UPDATE room_inventory
SET
    date = $3,
    total_inventory = $4,
    total_reserved = $5
WHERE hotel_id = $1 and room_id = $2
RETURNING *;

-- name: ListRoomInventory :many
SELECT * FROM room_inventory
WHERE hotel_id = COALESCE(sqlc.narg(hotel_id), hotel_id) AND (room_id = COALESCE(sqlc.narg(room_id), room_id))
ORDER BY sqlc.arg(hotel_id)
LIMIT sqlc.arg(page)
OFFSET sqlc.arg(offset1);