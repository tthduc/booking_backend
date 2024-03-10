// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: rate.sql

package db

import (
	"context"
)

const createRate = `-- name: CreateRate :one
INSERT INTO rate (  hotel_id,
                    room_id,
                    user_id,
                    rate) VALUES (
    $1, $2, $3, $4
) RETURNING hotel_id, room_id, user_id, rate, created_at
`

type CreateRateParams struct {
	HotelID int64 `json:"hotel_id"`
	RoomID  int64 `json:"room_id"`
	UserID  int64 `json:"user_id"`
	Rate    int64 `json:"rate"`
}

func (q *Queries) CreateRate(ctx context.Context, arg CreateRateParams) (Rate, error) {
	row := q.db.QueryRowContext(ctx, createRate,
		arg.HotelID,
		arg.RoomID,
		arg.UserID,
		arg.Rate,
	)
	var i Rate
	err := row.Scan(
		&i.HotelID,
		&i.RoomID,
		&i.UserID,
		&i.Rate,
		&i.CreatedAt,
	)
	return i, err
}