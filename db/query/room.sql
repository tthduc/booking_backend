-- name: CreateRoom :one
INSERT INTO room (
                  room_type_id,
                  hotel_id,
                  is_available,
                  status)
VALUES (
        $1, $2, $3, $4
        ) RETURNING *;

-- name: UpdateRoom :one
UPDATE room
SET room_type_id = $3
WHERE id = $1 AND hotel_id = $2
RETURNING *;

-- name: GetRoomByHotelId :one
SELECT * FROM room
WHERE hotel_id = $1;

-- name: DisableRoom :exec
UPDATE room
SET status = $2
WHERE id = $1
    RETURNING *;

-- name: GetRoom :one
SELECT * FROM room
WHERE id = $1 LIMIT 1;

-- name: DeleteRoom :exec
DELETE FROM room
WHERE id = $1;

-- name: ListRooms :many
SELECT * FROM room
ORDER BY id
    LIMIT $1
OFFSET $2;