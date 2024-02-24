-- name: CreateReservation :one
INSERT INTO reservation (
    hotel_id,
    room_id,
    start_date,
    end_date,
    status,
    amount,
    user_id)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetReservation :one
SELECT * FROM reservation
WHERE id = $1 FOR UPDATE;

-- name: ListReservations :many
SELECT * FROM reservation
WHERE
    hotel_id = $1 OR
    room_id = $2
ORDER BY id
    LIMIT $3
OFFSET $4;

-- name: UpdateReservation :one
UPDATE reservation
SET
    status = $3
WHERE hotel_id = $1 and room_id = $2
    RETURNING *;