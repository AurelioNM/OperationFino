package entity

import (
	"strings"
	"time"
)

type DateOnly struct {
	time.Time
}

const dateFormat = "2006-01-02"

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(dateFormat, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

type Customer struct {
	ID          *string    `json:"customer_id" db:"customer_id"`
	Name        string     `json:"name" db:"name"`
	Surname     string     `json:"surname" db:"surname"`
	Email       string     `json:"email" db:"email"`
	Birthdate   DateOnly  `json:"birthdate" db:"birthdate"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}
