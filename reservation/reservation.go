package reservation

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type ReservationComponent interface {
	List(ctx context.Context) ([]Reservation, error)
}

type reservationComponet struct {
	weaver.Implements[ReservationComponent]
	Reservations []Reservation `json:"reservations"`
}

type Reservation struct {
	ID          int
	FlightID    int
	UserID      int
	Status      string
	BookingDate string
	weaver.AutoMarshal
}

var dummyReservations = []Reservation{
	{ID: 1, FlightID: 1, UserID: 1, Status: "confirmed", BookingDate: "2023-01-01"},
	{ID: 2, FlightID: 2, UserID: 2, Status: "confirmed", BookingDate: "2023-01-02"},
	{ID: 3, FlightID: 3, UserID: 3, Status: "confirmed", BookingDate: "2023-01-03"},
	{ID: 4, FlightID: 4, UserID: 4, Status: "confirmed", BookingDate: "2023-01-04"},
	{ID: 5, FlightID: 5, UserID: 5, Status: "confirmed", BookingDate: "2023-01-05"},
}

func (r *reservationComponet) List(_ context.Context) ([]Reservation, error) {
	return dummyReservations, nil
}
