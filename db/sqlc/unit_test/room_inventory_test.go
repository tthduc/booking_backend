package unit_test

import (
	db "booking-backed/db/sqlc"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomRoomInventory(t *testing.T, hotel db.Hotel, roomType db.RoomType) db.RoomInventory {
	arg := db.CreateRoomInventoryParams{
		HotelID:        hotel.ID,
		RoomTypeID:     roomType.ID,
		TotalInventory: 100,
		TotalReserved:  0,
	}

	roomInventory, err := testQueries.CreateRoomInventory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, roomInventory)
	require.Equal(t, arg.HotelID, roomInventory.HotelID)
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

	roomType := createRandomRoomType(t)
	require.NotEmpty(t, hotel)

	createRandomRoomInventory(t, hotel, roomType)
}

func TestListRoomInventory(t *testing.T) {
	hotel := createRandomHotel(t)
	require.NotEmpty(t, hotel)

	roomType := createRandomRoomType(t)
	require.NotEmpty(t, hotel)

	for i := 0; i < 10; i++ {
		createRandomRoomInventory(t, hotel, roomType)
	}

	arg := db.ListRoomInventoryParams{
		Limit:   5,
		Offset:  5,
		HotelID: hotel.ID,
	}

	roomInventories, err := testQueries.ListRoomInventory(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, roomInventories, 5)

	for _, roomInventory := range roomInventories {
		require.NotEmpty(t, roomInventory)
	}
}
