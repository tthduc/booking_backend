package db

import (
	"booking-backed/enums"
	"context"
	"fmt"
	"time"
)

type ReserveTxParams struct {
	HotelID   int64     `json:"hotel_id"`
	RoomID    int64     `json:"room_id"`
	Amount    int64     `json:"amount"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    int32     `json:"status"`
	UserID    int64     `json:"user_id"`
}

type ReserveTxResult struct {
	Reservation Reservation `json:"reserve"`
}

var TxKey = struct{}{}

func (store *Store) ReserveTx(ctx context.Context, arg ReserveTxParams) (ReserveTxResult, error) {
	var result ReserveTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(TxKey)

		hotel, err := q.GetHotel(ctx, arg.HotelID)
		if err != nil {
			return err
		}
		fmt.Println(txName, "Get hotel: ", hotel)

		room, err := q.GetRoom(ctx, arg.RoomID)
		if err != nil {
			return err
		}
		fmt.Println(txName, "Get room: ", room)

		user, err := q.GetUser(ctx, arg.UserID)
		if err != nil {
			return err
		}
		fmt.Println(txName, "Get user: ", user)

		result.Reservation, err = q.CreateReservation(ctx, CreateReservationParams{
			HotelID:   hotel.ID,
			RoomID:    room.ID,
			Amount:    arg.Amount,
			StartDate: arg.StartDate,
			EndDate:   arg.EndDate,
			Status:    arg.Status,
			UserID:    user.ID,
		})
		if err != nil {
			return err
		}
		fmt.Println(txName, "Create reservation: ", result.Reservation)

		reservation, err := q.GetReservation(ctx, result.Reservation.ID)
		if err != nil {
			return err
		}
		fmt.Println(txName, "Get reservation: ", reservation)

		_, err = q.UpdateReservation(ctx, UpdateReservationParams{
			HotelID: reservation.HotelID,
			RoomID:  reservation.RoomID,
			Status:  enums.Approve,
		})
		if err != nil {
			return err
		}
		fmt.Println(txName, "Update reservation")

		roomInventory, err := q.GetRoomInventory(ctx, GetRoomInventoryParams{
			HotelID:    reservation.HotelID,
			RoomTypeID: room.RoomTypeID,
		})
		if err != nil {
			return err
		}
		fmt.Println(txName, "Get roomInventory: ", roomInventory)

		_, err = q.UpdateRoomInventory(ctx, UpdateRoomInventoryParams{
			HotelID:    reservation.HotelID,
			RoomTypeID: room.RoomTypeID,
			Date:       time.Now(),
		})
		if err != nil {
			return err
		}
		fmt.Println(txName, "Update room inventory")

		return err
	})

	return result, err
}
