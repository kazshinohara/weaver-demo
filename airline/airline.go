package airline

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type AirlineComponent interface {
	List(ctx context.Context) ([]Airline, error)
}

type airlineComponent struct {
	weaver.Implements[AirlineComponent]
	Airlines []Airline `json:"airlines"`
}

type Airline struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	weaver.AutoMarshal
}

var dummyAirlines = []Airline{
	{ID: 1, Name: "Airline A"},
	{ID: 2, Name: "Airline B"},
	{ID: 3, Name: "Airline C"},
	{ID: 4, Name: "Airline D"},
	{ID: 5, Name: "Airline E"},
}

func (a *airlineComponent) List(_ context.Context) ([]Airline, error) {
	return dummyAirlines, nil
}
