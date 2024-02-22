-- name: CreateRoom :one
INSERT INTO room (
                  room_type_id,
                  hotel_id,
                  is_available)
VALUES (
        $1, $2, $3
        ) RETURNING *;

-- name: UpdateRoom :exec
UPDATE room
SET room_type_id = $3
WHERE id = $1 AND hotel_id = $2
RETURNING *;

-- name: GetRoomByHotelId :one
SELECT * FROM room
WHERE $1;

-- name: DisableRoom :exec
UPDATE room
SET status = $2
WHERE id = $1
    RETURNING *;