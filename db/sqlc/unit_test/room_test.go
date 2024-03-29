package unit_test

import (
	db "booking-backed/db/sqlc"
	"booking-backed/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomRoom(t *testing.T, hotel db.Hotel, roomType db.RoomType) db.Room {
	arg := db.CreateRoomParams{
		RoomTypeID:  roomType.ID,
		HotelID:     hotel.ID,
		Name:        util.RandomString(10),
		IsAvailable: 1,
		Status:      1,
	}

	room, err := testQueries.CreateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room)
	require.Equal(t, arg.Name, room.Name)
	require.Equal(t, arg.RoomTypeID, room.RoomTypeID)
	require.Equal(t, arg.HotelID, room.HotelID)
	require.Equal(t, arg.IsAvailable, room.IsAvailable)
	require.Equal(t, arg.Status, room.Status)

	require.NotZero(t, room.ID)
	require.NotZero(t, room.CreatedAt)

	return room
}

func TestCreateRoom(t *testing.T) {
	hotel := createRandomHotel(t)
	require.NotEmpty(t, hotel)

	roomType := createRandomRoomType(t)
	require.NotEmpty(t, roomType)

	createRandomRoom(t, hotel, roomType)
}

func TestGetRoom(t *testing.T) {
	hotel := createRandomHotel(t)
	require.NotEmpty(t, hotel)

	roomType := createRandomRoomType(t)
	require.NotEmpty(t, roomType)

	room1 := createRandomRoom(t, hotel, roomType)
	room2, err := testQueries.GetRoom(context.Background(), room1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, room2)

	require.Equal(t, room1.ID, room2.ID)
	require.Equal(t, room1.RoomTypeID, room2.RoomTypeID)
	require.Equal(t, room1.HotelID, room2.HotelID)
	require.Equal(t, room1.IsAvailable, room2.IsAvailable)
	require.Equal(t, room1.CreatedAt, room2.CreatedAt)
}

func TestGetRoomByHotelId(t *testing.T) {
	hotel := createRandomHotel(t)
	require.NotEmpty(t, hotel)

	roomType := createRandomRoomType(t)
	require.NotEmpty(t, roomType)

	room1 := createRandomRoom(t, hotel, roomType)
	room2, err := testQueries.GetRoomByHotelId(context.Background(), room1.HotelID)

	require.NoError(t, err)
	require.NotEmpty(t, room2)

	require.Equal(t, room1.ID, room2.ID)
	require.Equal(t, room1.RoomTypeID, room2.RoomTypeID)
	require.Equal(t, room1.HotelID, room2.HotelID)
	require.Equal(t, room1.IsAvailable, room2.IsAvailable)
	require.Equal(t, room1.CreatedAt, room2.CreatedAt)
}

func TestUpdateRoom(t *testing.T) {
	hotel := createRandomHotel(t)
	require.NotEmpty(t, hotel)

	roomType := createRandomRoomType(t)
	require.NotEmpty(t, roomType)

	room1 := createRandomRoom(t, hotel, roomType)
	require.NotEmpty(t, room1)

	hotel, err := testQueries.GetHotel(context.Background(), room1.HotelID)
	require.NoError(t, err)
	require.NotEmpty(t, hotel)

	arg := db.UpdateRoomParams{
		ID:         room1.ID,
		HotelID:    hotel.ID,
		RoomTypeID: roomType.ID,
		Name:       util.RandomString(10),
	}

	room2, err := testQueries.UpdateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room2)

	require.Equal(t, room1.ID, room2.ID)
	require.Equal(t, arg.RoomTypeID, room2.RoomTypeID)
	require.Equal(t, room1.CreatedAt, room2.CreatedAt)
	require.WithinDuration(t, room1.CreatedAt, room2.CreatedAt, time.Second)
}

func TestDeleteRoom(t *testing.T) {
	hotel := createRandomHotel(t)
	require.NotEmpty(t, hotel)

	roomType := createRandomRoomType(t)
	require.NotEmpty(t, roomType)

	room1 := createRandomRoom(t, hotel, roomType)
	err := testQueries.DeleteRoom(context.Background(), room1.ID)
	require.NoError(t, err)

	hotel2, err := testQueries.GetRoom(context.Background(), room1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, hotel2)
}

func TestListRoom(t *testing.T) {
	hotel := createRandomHotel(t)
	require.NotEmpty(t, hotel)

	roomType := createRandomRoomType(t)
	require.NotEmpty(t, roomType)

	for i := 0; i < 10; i++ {
		createRandomRoom(t, hotel, roomType)
	}

	arg := db.ListRoomsParams{
		Limit:  5,
		Offset: 5,
	}

	rooms, err := testQueries.ListRooms(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, rooms, 5)

	for _, room := range rooms {
		require.NotEmpty(t, room)
	}
}
