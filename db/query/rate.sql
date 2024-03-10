-- name: CreateRate :one
INSERT INTO rate (  hotel_id,
                    room_id,
                    user_id,
                    rate) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateRate :one
UPDATE rate
SET rate = $4
WHERE hotel_id = $1 AND room_id = $2 AND user_id = $3
RETURNING *;

