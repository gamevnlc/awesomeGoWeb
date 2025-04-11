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
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var newID int
	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	stmt := `insert into reservations (first_name, last_name, email, phone, 
                          start_date, end_date, room_id, created_at, updated_at)
                          values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// InsertRoomRestriction insert a room restriction into database
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	stmt := `insert into room_restrictions (start_date, end_date, room_id, reservation_id, restriction_id, 
                               created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		r.RestrictionID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}

// SearchAvailabilityByDates return true if availability exists for room ID and false if no availability exists
func (m *postgresDBRepo) SearchAvailabilityByDates(start time.Time, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	var numRows int
	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `
		select
			count(id)
		from
		    room_restrictions
		where
		    room_id = $1
		    $2 < end_date and $3 > start_date;
	`

	row := m.DB.QueryRowContext(ctx, query, roomID, start, end)
	err := row.Scan(&numRows)

	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}
	return false, nil
}
