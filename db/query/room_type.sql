-- name: CreateRoomType :one
INSERT INTO room_type (
                       name
                       ) VALUES (
    $1
) RETURNING *;

-- name: GetRoomType :one
SELECT * FROM room_type
WHERE id = $1 LIMIT 1;

-- name: ListRoomType :many
SELECT * FROM room_type
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: UpdateRoomType :one
UPDATE room_type
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteRoomType :exec
DELETE FROM room_type
WHERE id = $1;