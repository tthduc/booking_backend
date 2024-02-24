// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: room_inventory.sql

package db

import (
	"context"
	"time"
)

const createRoomInventory = `-- name: CreateRoomInventory :one
INSERT INTO room_inventory (hotel_id,
                   room_type_id,
                   date,
                   total_inventory,
                   total_reserved) VALUES (
    $1, $2, $3, $4, $5
) RETURNING hotel_id, room_type_id, date, total_inventory, total_reserved, created_at
`

type CreateRoomInventoryParams struct {
	HotelID        int64     `json:"hotel_id"`
	RoomTypeID     int64     `json:"room_type_id"`
	Date           time.Time `json:"date"`
	TotalInventory int32     `json:"total_inventory"`
	TotalReserved  int32     `json:"total_reserved"`
}

func (q *Queries) CreateRoomInventory(ctx context.Context, arg CreateRoomInventoryParams) (RoomInventory, error) {
	row := q.db.QueryRowContext(ctx, createRoomInventory,
		arg.HotelID,
		arg.RoomTypeID,
		arg.Date,
		arg.TotalInventory,
		arg.TotalReserved,
	)
	var i RoomInventory
	err := row.Scan(
		&i.HotelID,
		&i.RoomTypeID,
		&i.Date,
		&i.TotalInventory,
		&i.TotalReserved,
		&i.CreatedAt,
	)
	return i, err
}

const getRoomInventory = `-- name: GetRoomInventory :one
SELECT hotel_id, room_type_id, date, total_inventory, total_reserved, created_at FROM room_inventory
WHERE hotel_id = $1 AND room_type_id = $2
`

type GetRoomInventoryParams struct {
	HotelID    int64 `json:"hotel_id"`
	RoomTypeID int64 `json:"room_type_id"`
}

func (q *Queries) GetRoomInventory(ctx context.Context, arg GetRoomInventoryParams) (RoomInventory, error) {
	row := q.db.QueryRowContext(ctx, getRoomInventory, arg.HotelID, arg.RoomTypeID)
	var i RoomInventory
	err := row.Scan(
		&i.HotelID,
		&i.RoomTypeID,
		&i.Date,
		&i.TotalInventory,
		&i.TotalReserved,
		&i.CreatedAt,
	)
	return i, err
}

const listRoomInventory = `-- name: ListRoomInventory :many
SELECT hotel_id, room_type_id, date, total_inventory, total_reserved, created_at FROM room_inventory
WHERE hotel_id = $3
ORDER BY hotel_id
LIMIT $1
OFFSET $2
`

type ListRoomInventoryParams struct {
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
	HotelID int64 `json:"hotel_id"`
}

func (q *Queries) ListRoomInventory(ctx context.Context, arg ListRoomInventoryParams) ([]RoomInventory, error) {
	rows, err := q.db.QueryContext(ctx, listRoomInventory, arg.Limit, arg.Offset, arg.HotelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RoomInventory
	for rows.Next() {
		var i RoomInventory
		if err := rows.Scan(
			&i.HotelID,
			&i.RoomTypeID,
			&i.Date,
			&i.TotalInventory,
			&i.TotalReserved,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRoomInventory = `-- name: UpdateRoomInventory :one
UPDATE room_inventory
SET
    date = $3,
    total_inventory = total_inventory - 1,
    total_reserved = total_reserved + 1
WHERE hotel_id = $1 and room_type_id = $2
RETURNING hotel_id, room_type_id, date, total_inventory, total_reserved, created_at
`

type UpdateRoomInventoryParams struct {
	HotelID    int64     `json:"hotel_id"`
	RoomTypeID int64     `json:"room_type_id"`
	Date       time.Time `json:"date"`
}

func (q *Queries) UpdateRoomInventory(ctx context.Context, arg UpdateRoomInventoryParams) (RoomInventory, error) {
	row := q.db.QueryRowContext(ctx, updateRoomInventory, arg.HotelID, arg.RoomTypeID, arg.Date)
	var i RoomInventory
	err := row.Scan(
		&i.HotelID,
		&i.RoomTypeID,
		&i.Date,
		&i.TotalInventory,
		&i.TotalReserved,
		&i.CreatedAt,
	)
	return i, err
}
