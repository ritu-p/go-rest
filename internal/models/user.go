package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

var users = []User{
	{ID: uuid.New(), Name: "Alice", Email: "test@email"},
	{ID: uuid.New(), Name: "Bob", Email: "test2@email"},
	{ID: uuid.New(), Name: "Charlie", Email: "test3@email"},
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *User) BeforeCreate(tx any) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
