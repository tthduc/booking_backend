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

func createRandomRoomType(t *testing.T) db.RoomType {
	name := util.RandomString(10)
	roomType, err := testQueries.CreateRoomType(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, roomType)

	require.Equal(t, name, roomType.Name)

	require.NotZero(t, roomType.ID)
	require.NotZero(t, roomType.CreatedAt)

	return roomType
}

func TestCreateRoomType(t *testing.T) {
	createRandomRoomType(t)
}

func TestGetRoomType(t *testing.T) {
	roomType1 := createRandomRoomType(t)
	roomType2, err := testQueries.GetRoomType(context.Background(), roomType1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, roomType2)

	require.Equal(t, roomType1.ID, roomType2.ID)
	require.Equal(t, roomType1.Name, roomType2.Name)
	require.Equal(t, roomType1.CreatedAt, roomType2.CreatedAt)
}

func TestUpdateRoomType(t *testing.T) {
	roomType1 := createRandomRoomType(t)

	arg := db.UpdateRoomTypeParams{
		ID:   roomType1.ID,
		Name: util.RandomString(10),
	}

	roomType2, err := testQueries.UpdateRoomType(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, roomType2)

	require.Equal(t, roomType1.ID, roomType2.ID)
	require.Equal(t, arg.Name, roomType2.Name)
	require.Equal(t, roomType1.CreatedAt, roomType2.CreatedAt)
	require.WithinDuration(t, roomType1.CreatedAt, roomType2.CreatedAt, time.Second)
}

func TestDeleteRoomType(t *testing.T) {
	roomType1 := createRandomRoomType(t)
	err := testQueries.DeleteRoomType(context.Background(), roomType1.ID)
	require.NoError(t, err)

	roomType2, err := testQueries.GetRoomType(context.Background(), roomType1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, roomType2)
}

func TestListRoomType(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomRoomType(t)
	}

	arg := db.ListRoomTypeParams{
		Limit:  5,
		Offset: 5,
	}

	roomTypes, err := testQueries.ListRoomType(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, roomTypes, 5)

	for _, roomType := range roomTypes {
		require.NotEmpty(t, roomType)
	}
}
