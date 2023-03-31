package user

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type UserComponent interface {
	List(ctx context.Context) ([]User, error)
}

type userComponent struct {
	weaver.Implements[UserComponent]
	Users []User `json:"users"`
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	weaver.AutoMarshal
}

var dummyUsers = []User{
	{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com", Phone: "123-456-7890"},
	{ID: 2, FirstName: "Jane", LastName: "Doe", Email: "jane@example.com", Phone: "234-567-8901"},
	{ID: 3, FirstName: "Bob", LastName: "Smith", Email: "bob@example.com", Phone: "345-678-9012"},
	{ID: 4, FirstName: "Alice", LastName: "Johnson", Email: "alice@example.com", Phone: "456-789-0123"},
	{ID: 5, FirstName: "Charlie", LastName: "Patel", Email: "charlie@example.com", Phone: "987-654-3210"},
}

func (u *userComponent) List(_ context.Context) ([]User, error) {
	return dummyUsers, nil
}