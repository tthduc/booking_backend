-- name: CreateRate :one
INSERT INTO rate (  hotel_id,
                    room_id,
                    user_id,
                    rate) VALUES (
    $1, $2, $3, $4
) RETURNING *;