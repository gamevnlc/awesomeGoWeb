package repository

import (
	"awesomeWeb/internal/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start time.Time, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start time.Time, end time.Time) ([]models.Room, error)
}
