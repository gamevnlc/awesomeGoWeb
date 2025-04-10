package dbrepo

import (
	"awesomeWeb/internal/models"
	"context"
	"time"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation insert reservation into database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	stmt := `insert into reservations (first_name, last_name, email, phone, 
                          start_date, end_date, room_id, created_at, update_at)
                          values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := m.DB.ExecContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
