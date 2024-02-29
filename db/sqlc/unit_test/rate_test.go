package unit_test

import (
	db "booking-backed/db/sqlc"
	"booking-backed/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
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

func TestUpdateRate(t *testing.T) {
	rate1 := createRandomRate(t)
	arg := db.UpdateRateParams{
		HotelID: rate1.HotelID,
		RoomID:  rate1.RoomID,
		UserID:  rate1.UserID,
		Rate:    util.RandomInt(0, 5),
	}
	rate2, err := testQueries.UpdateRate(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, rate2)
	require.Equal(t, rate1.HotelID, rate2.HotelID)
	require.Equal(t, rate1.RoomID, rate2.RoomID)
	require.Equal(t, rate1.UserID, rate2.UserID)
	require.Equal(t, arg.Rate, rate2.Rate)
	require.WithinDuration(t, rate1.CreatedAt, rate2.CreatedAt, time.Second)
}
