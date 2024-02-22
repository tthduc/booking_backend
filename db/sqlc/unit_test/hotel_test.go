package unit_test

import (
	"booking-backed/db/sqlc"
	"booking-backed/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomHotel(t *testing.T) db.Hotel {
	arg := db.CreateHotelParams{
		Name:     util.RandomOwner(),
		Address:  util.RandomAddress(),
		Location: util.RandomAddress(),
	}

	hotel, err := testQueries.CreateHotel(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, hotel)
	require.Equal(t, arg.Name, hotel.Name)
	require.Equal(t, arg.Address, hotel.Address)
	require.Equal(t, arg.Location, hotel.Location)

	require.NotZero(t, hotel.ID)
	require.NotZero(t, hotel.CreatedAt)

	return hotel
}

func TestCreateHotel(t *testing.T) {
	createRandomHotel(t)
}

func TestGetHotel(t *testing.T) {
	hotel1 := createRandomHotel(t)
	hotel2, err := testQueries.GetHotel(context.Background(), hotel1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, hotel2)

	require.Equal(t, hotel1.ID, hotel2.ID)
	require.Equal(t, hotel1.Name, hotel2.Name)
	require.Equal(t, hotel1.Address, hotel2.Address)
	require.Equal(t, hotel1.Location, hotel2.Location)
	require.Equal(t, hotel1.CreatedAt, hotel2.CreatedAt)
}

func TestUpdateHotel(t *testing.T) {
	hotel1 := createRandomHotel(t)

	arg := db.UpdateHotelParams{
		ID:       hotel1.ID,
		Name:     util.RandomOwner(),
		Address:  util.RandomAddress(),
		Location: util.RandomAddress(),
	}

	hotel2, err := testQueries.UpdateHotel(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, hotel2)
	require.Equal(t, hotel1.ID, hotel2.ID)
	require.Equal(t, arg.Name, hotel2.Name)
	require.Equal(t, arg.Address, hotel2.Address)
	require.Equal(t, arg.Location, hotel2.Location)
	require.Equal(t, hotel1.CreatedAt, hotel2.CreatedAt)
	require.WithinDuration(t, hotel1.CreatedAt, hotel2.CreatedAt, time.Second)
}

func TestDeleteHotel(t *testing.T) {
	hotel1 := createRandomHotel(t)
	err := testQueries.DeleteHotel(context.Background(), hotel1.ID)
	require.NoError(t, err)

	hotel2, err := testQueries.GetHotel(context.Background(), hotel1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, hotel2)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomHotel(t)
	}

	arg := db.ListHotelsParams{
		Limit:  5,
		Offset: 5,
	}

	hotels, err := testQueries.ListHotels(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, hotels, 5)

	for _, hotel := range hotels {
		require.NotEmpty(t, hotel)
	}
}
