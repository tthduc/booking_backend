package unit_test

import (
	db "booking-backed/db/sqlc"
	"booking-backed/enums"
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestReserveTx(t *testing.T) {
	store := db.NewStore(testDB)

	roomType := createRandomRoomType(t)
	require.NotEmpty(t, roomType)

	room := createRandomRoom(t, roomType)
	require.NotEmpty(t, room)

	hotel := createRandomHotel(t)
	require.NotEmpty(t, hotel)

	user := createRandomUser(t)
	require.NotEmpty(t, user)

	roomInventory := createRandomRoomInventory(t, hotel, roomType)
	require.NotEmpty(t, roomInventory)

	n := 3
	amount := int64(100)

	// run n concurrent reservation transaction
	errs := make(chan error)
	results := make(chan db.ReserveTxResult)

	// run n concurrent reservation transaction
	for i := 0; i < n; i++ {
		//txName := make(chan db.ReserveTxResult)
		txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			ctx := context.WithValue(context.Background(), db.TxKey, txName)

			result, err := store.ReserveTx(ctx, db.ReserveTxParams{
				HotelID:   hotel.ID,
				RoomID:    room.ID,
				Amount:    amount,
				UserID:    user.ID,
				StartDate: time.Now(),
				EndDate:   time.Now().AddDate(0, 0, 3),
				Status:    enums.Waiting,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		if err != nil && err.Error() == db.RoomInventoryError {
			require.Equal(t, err.Error(), db.RoomInventoryError)
		} else {
			require.NoError(t, err)
		}

		result := <-results
		require.NotEmpty(t, result)

		// check reservation
		reservation := result.Reservation
		require.NotEmpty(t, reservation)
		require.Equal(t, hotel.ID, reservation.HotelID)
		require.Equal(t, room.ID, reservation.RoomID)
		require.Equal(t, amount, reservation.Amount)
		require.NotZero(t, reservation.ID)
		require.NotZero(t, reservation.CreatedAt)
	}
}
