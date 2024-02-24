package unit_test

import (
	db "booking-backed/db/sqlc"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomRoomInventory(t *testing.T, hotel db.Hotel) db.RoomInventory {
	room := createRandomRoom(t)
	require.NotEmpty(t, room)

	arg := db.CreateRoomInventoryParams{
		HotelID:        hotel.ID,
		RoomID:         room.ID,
		RoomTypeID:     room.RoomTypeID,
		TotalInventory: 100,
		TotalReserved:  0,
	}

	roomInventory, err := testQueries.CreateRoomInventory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, roomInventory)
	require.Equal(t, arg.HotelID, roomInventory.HotelID)
	require.Equal(t, arg.RoomID, roomInventory.RoomID)
	require.Equal(t, arg.RoomTypeID, roomInventory.RoomTypeID)
	require.Equal(t, arg.TotalInventory, roomInventory.TotalInventory)
	require.Equal(t, arg.TotalReserved, roomInventory.TotalReserved)

	require.NotZero(t, hotel.ID)
	require.NotZero(t, hotel.CreatedAt)

	return roomInventory
}

func TestCreateRoomInventory(t *testing.T) {
	hotel := createRandomHotel(t)
	require.NotEmpty(t, hotel)
	createRandomRoomInventory(t, hotel)
}

func TestListRoomInventory(t *testing.T) {
	hotel := createRandomHotel(t)
	require.NotEmpty(t, hotel)

	for i := 0; i < 10; i++ {
		createRandomRoomInventory(t, hotel)
	}

	arg := db.ListRoomInventoryParams{
		Page:    5,
		Offset1: 5,
		//HotelID: sql.NullInt64{
		//	Int64: 34,
		//	Valid: true,
		//},
		//RoomID: sql.NullInt64{
		//	Int64: 26,
		//	Valid: true,
		//},
	}

	roomInventories, err := testQueries.ListRoomInventory(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, roomInventories, 5)

	for _, roomInventory := range roomInventories {
		require.NotEmpty(t, roomInventory)
	}
}
