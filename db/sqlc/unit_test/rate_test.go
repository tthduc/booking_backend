package unit_test

import (
	db "booking-backed/db/sqlc"
	"booking-backed/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomRate(t *testing.T) db.Rate {
	hotel := createRandomHotel(t)
	require.NotEmpty(t, hotel)

	roomType := createRandomRoomType(t)
	require.NotEmpty(t, roomType)

	room := createRandomRoom(t, hotel, roomType)
	require.NotEmpty(t, room)

	user := createRandomUser(t)
	require.NotEmpty(t, user)

	arg := db.CreateRateParams{
		HotelID: hotel.ID,
		RoomID:  room.ID,
		UserID:  user.ID,
		Rate:    util.RandomInt(0, 5),
	}

	rate, err := testQueries.CreateRate(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, rate)
	require.Equal(t, arg.HotelID, rate.HotelID)
	require.Equal(t, arg.RoomID, rate.RoomID)
	require.Equal(t, arg.Rate, rate.Rate)
	require.Equal(t, arg.UserID, rate.UserID)

	require.NotZero(t, rate.CreatedAt)

	return rate
}

func TestCreateRate(t *testing.T) {
	createRandomRate(t)
}
