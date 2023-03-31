package flight

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type FlightComponent interface {
	List(ctx context.Context) ([]Flight, error)
}

type flightComponent struct {
	weaver.Implements[FlightComponent]
	Flights []Flight `json:"flights"`
}

type Flight struct {
	ID             int
	Origin         string
	Destination    string
	AirlineID      int
	Departure      string
	Arrival        string
	AvailableSeats int
	weaver.AutoMarshal
}

var dummyFlights = []Flight{
	{ID: 1, Origin: "Tokyo", Destination: "Osaka", AirlineID: 1, Departure: "10:00", Arrival: "11:30", AvailableSeats: 10},
	{ID: 2, Origin: "Tokyo", Destination: "Osaka", AirlineID: 2, Departure: "12:00", Arrival: "13:30", AvailableSeats: 20},
	{ID: 3, Origin: "Tokyo", Destination: "Osaka", AirlineID: 3, Departure: "14:00", Arrival: "15:30", AvailableSeats: 30},
	{ID: 4, Origin: "Tokyo", Destination: "Osaka", AirlineID: 4, Departure: "16:00", Arrival: "17:30", AvailableSeats: 40},
	{ID: 5, Origin: "Tokyo", Destination: "Osaka", AirlineID: 5, Departure: "18:00", Arrival: "19:30", AvailableSeats: 50},
}

func (f *flightComponent) List(_ context.Context) ([]Flight, error) {
	return dummyFlights, nil
}