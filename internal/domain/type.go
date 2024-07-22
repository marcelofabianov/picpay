package domain

import (
	"time"

	"github.com/google/uuid"
)

type ID string
type Email string
type Password string
type DocumentRegistry string
type Amount float64
type CreatedAt time.Time
type UpdatedAt time.Time
type DeletedAt time.Time
type Version int

func NewID() ID {
	return ID(uuid.New().String())
}

func CreatedAtNow() CreatedAt {
	return CreatedAt(time.Now())
}

func UpdatedAtNow() UpdatedAt {
	return UpdatedAt(time.Now())
}

func InitVersion() Version {
	return Version(1)
}

func (c CreatedAt) Format() string {
	return time.Time(c).Format(time.RFC3339)
}

func (u UpdatedAt) Format() string {
	return time.Time(u).Format(time.RFC3339)
}
