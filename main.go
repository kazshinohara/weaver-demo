package main

import (
	"context"
	"encoding/json"
	"flight-master/airline"
	"flight-master/flight"
	"flight-master/reservation"
	"flight-master/user"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	ctx := context.Background()
	root := weaver.Init(ctx)
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := root.Listener("weaver-demo", opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening on %v\n", lis)

	airlineComponent, err := weaver.Get[airline.AirlineComponent](root)
	flightComponent, err := weaver.Get[flight.FlightComponent](root)
	reservationComponent, err := weaver.Get[reservation.ReservationComponent](root)
	userComponent, err := weaver.Get[user.UserComponent](root)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/airlines", func(w http.ResponseWriter, r *http.Request) {
		airlines, err := airlineComponent.List(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonRes, err := json.Marshal(airlines)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
	})

	r.Get("/flights", func(w http.ResponseWriter, r *http.Request) {
		flights, err := flightComponent.List(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonRes, err := json.Marshal(flights)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
	})

	r.Get("/reservations", func(w http.ResponseWriter, r *http.Request) {
		reservations, err := reservationComponent.List(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonRes, err := json.Marshal(reservations)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := userComponent.List(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonRes, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
	})

	http.Serve(lis, r)
}
