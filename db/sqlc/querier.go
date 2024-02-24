// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateHotel(ctx context.Context, arg CreateHotelParams) (Hotel, error)
	CreateReservation(ctx context.Context, arg CreateReservationParams) (Reservation, error)
	CreateRoom(ctx context.Context, arg CreateRoomParams) (Room, error)
	CreateRoomInventory(ctx context.Context, arg CreateRoomInventoryParams) (RoomInventory, error)
	CreateRoomType(ctx context.Context, name string) (RoomType, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteHotel(ctx context.Context, id int64) error
	DeleteRoom(ctx context.Context, id int64) error
	DeleteRoomType(ctx context.Context, id int64) error
	DisableRoom(ctx context.Context, arg DisableRoomParams) error
	GetHotel(ctx context.Context, id int64) (Hotel, error)
	GetReservation(ctx context.Context, id int64) (Reservation, error)
	GetRoom(ctx context.Context, id int64) (Room, error)
	GetRoomByHotelId(ctx context.Context, hotelID int64) (Room, error)
	GetRoomInventory(ctx context.Context, arg GetRoomInventoryParams) (RoomInventory, error)
	GetRoomType(ctx context.Context, id int64) (RoomType, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	ListHotels(ctx context.Context, arg ListHotelsParams) ([]Hotel, error)
	ListReservations(ctx context.Context, arg ListReservationsParams) ([]Reservation, error)
	ListRoomInventory(ctx context.Context, arg ListRoomInventoryParams) ([]RoomInventory, error)
	ListRoomType(ctx context.Context, arg ListRoomTypeParams) ([]RoomType, error)
	ListRooms(ctx context.Context, arg ListRoomsParams) ([]Room, error)
	UpdateHotel(ctx context.Context, arg UpdateHotelParams) (Hotel, error)
	UpdateReservation(ctx context.Context, arg UpdateReservationParams) (Reservation, error)
	UpdateRoom(ctx context.Context, arg UpdateRoomParams) (Room, error)
	UpdateRoomInventory(ctx context.Context, arg UpdateRoomInventoryParams) (RoomInventory, error)
	UpdateRoomType(ctx context.Context, arg UpdateRoomTypeParams) (RoomType, error)
}

var _ Querier = (*Queries)(nil)
