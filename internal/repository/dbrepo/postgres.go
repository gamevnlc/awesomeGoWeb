package dbrepo

import (
	"awesomeWeb/internal/models"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
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

// SearchAvailabilityByDatesByRoomID return true if availability exists for room ID and false if no availability exists
func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomID(start time.Time, end time.Time, roomID int) (bool, error) {
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
		    room_id = $1 and
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

// SearchAvailabilityForAllRooms return a slice of available rooms if any for give date range
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start time.Time, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var rooms []models.Room
	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `
		select
			r.id, r.room_name
		from
		    rooms r
		where r.id not in (
		    select room_id from room_restrictions rr where $1 < rr.end_date and $2 > rr.start_date
		)
	`

	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
	}
	defer rows.Close()
	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	if err := rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

// GetRoomByID get a room by id
func (m *postgresDBRepo) GetRoomByID(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	var room models.Room
	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `
		select id, room_name, created_at, updated_at
		from
		    rooms
		where id = $1
	`
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	if err != nil {
		return room, err
	}
	return room, nil
}

// GetUserByID return user by id
func (m *postgresDBRepo) GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var user models.User
	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `
		select id, first_name, last_name, email, password, access_level, created_at, updated_at
		from users where id = $1
	`

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.AccessLevel,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}
	return user, nil
}

// UpdateUser update a user in a database
func (m *postgresDBRepo) UpdateUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `
		update users set first_name = $1, last_name = $2, email = $3, access_level = $4, updated_at = $5
	`

	_, err := m.DB.ExecContext(
		ctx, query,
		u.FirstName,
		u.LastName,
		u.Email,
		u.AccessLevel,
		time.Now(),
	)

	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo) Authenticate(email string, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var id int
	var hashedPassword string

	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `select id, password from users where email = $1`
	row := m.DB.QueryRowContext(ctx, query, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("wrong password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}

// AllReservations return all reservations
func (m *postgresDBRepo) AllReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	var reservations []models.Reservation

	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `
		select r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date, r.end_date, r.room_id, r.created_at,
		r.updated_at, r.processed, rm.id, rm.room_name
		from reservations r
		left join rooms rm on r.room_id = rm.id
		order by r.start_date asc
	`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return reservations, err
	}
	defer rows.Close()
	for rows.Next() {
		var reservation models.Reservation
		err := rows.Scan(
			&reservation.ID,
			&reservation.FirstName,
			&reservation.LastName,
			&reservation.Email,
			&reservation.Phone,
			&reservation.StartDate,
			&reservation.EndDate,
			&reservation.RoomID,
			&reservation.CreatedAt,
			&reservation.UpdatedAt,
			&reservation.Processed,
			&reservation.Room.ID,
			&reservation.Room.RoomName,
		)
		if err != nil {
			return reservations, err
		}

		reservations = append(reservations, reservation)
	}
	if err := rows.Err(); err != nil {
		return reservations, err
	}
	return reservations, nil
}

func (m *postgresDBRepo) AllNewReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	var reservations []models.Reservation

	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `
		select r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date, r.end_date, r.room_id, r.created_at,
		r.updated_at, rm.id, rm.room_name
		from reservations r
		left join rooms rm on r.room_id = rm.id
		where processed = 0
		order by r.start_date asc
	`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return reservations, err
	}
	defer rows.Close()
	for rows.Next() {
		var reservation models.Reservation
		err := rows.Scan(
			&reservation.ID,
			&reservation.FirstName,
			&reservation.LastName,
			&reservation.Email,
			&reservation.Phone,
			&reservation.StartDate,
			&reservation.EndDate,
			&reservation.RoomID,
			&reservation.CreatedAt,
			&reservation.UpdatedAt,
			&reservation.Room.ID,
			&reservation.Room.RoomName,
		)
		if err != nil {
			return reservations, err
		}

		reservations = append(reservations, reservation)
	}
	if err := rows.Err(); err != nil {
		return reservations, err
	}
	return reservations, nil
}

// GetReservationByID return one reservation by id
func (m *postgresDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	var reservation models.Reservation

	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `
		select r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date, r.end_date,
			r.room_id, r.created_at, r.updated_at, r.processed, rm.id, rm.room_name
		from reservations r
		left join rooms rm on r.room_id = rm.id
		where r.id = $1
	`
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&reservation.ID,
		&reservation.FirstName,
		&reservation.LastName,
		&reservation.Email,
		&reservation.Phone,
		&reservation.StartDate,
		&reservation.EndDate,
		&reservation.RoomID,
		&reservation.CreatedAt,
		&reservation.UpdatedAt,
		&reservation.Processed,
		&reservation.Room.ID,
		&reservation.Room.RoomName,
	)
	if err != nil {
		return reservation, err
	}
	return reservation, nil
}

func (m *postgresDBRepo) UpdateReservation(u models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `
		update reservations set first_name = $1, last_name = $2, email = $3, phone = $4, updated_at = $5
		where id = $6
	`

	_, err := m.DB.ExecContext(
		ctx, query,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Phone,
		time.Now(),
		u.ID,
	)

	if err != nil {
		return err
	}
	return nil
}

// DeleteReservation delete one reservation by id
func (m *postgresDBRepo) DeleteReservation(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `delete from reservations where id = $1`

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProcessedForReservation update processed for reservation
func (m *postgresDBRepo) UpdateProcessedForReservation(id, processed int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `update reservations set processed = $1 where id = $2`
	_, err := m.DB.ExecContext(ctx, query, processed, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo) AllRooms() ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	var rooms []models.Room
	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `select id , room_name, created_at, updated_at from rooms order by room_name`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return rooms, err
	}
	defer rows.Close()
	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		return rooms, err
	}
	return rooms, nil
}

// GetRestrictionsFroRoomByDate returns restriction for a room by date range
func (m *postgresDBRepo) GetRestrictionsFroRoomByDate(roomId int, start time.Time, end time.Time) ([]models.RoomRestriction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var restrictions []models.RoomRestriction

	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `
		select id, coalesce(reservation_id, 0), restriction_id, room_id, start_date, end_date
		from room_restrictions where $1 < end_date and $2 >= start_date and room_id = $3
	`
	rows, err := m.DB.QueryContext(ctx, query, start, end, roomId)
	if err != nil {
		return restrictions, err
	}
	defer rows.Close()
	for rows.Next() {
		var restriction models.RoomRestriction
		err := rows.Scan(
			&restriction.ID,
			&restriction.ReservationID,
			&restriction.RestrictionID,
			&restriction.RoomID,
			&restriction.StartDate,
			&restriction.EndDate,
		)
		if err != nil {
			return restrictions, err
		}
		restrictions = append(restrictions, restriction)
	}
	if err := rows.Err(); err != nil {
		return restrictions, err
	}
	return restrictions, nil
}

// InsertBlockForRoom insert a room restriction
func (m *postgresDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `
		insert into room_restrictions (start_date, end_date, room_id, restriction_id, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6)
	`

	_, err := m.DB.ExecContext(
		ctx, query,
		startDate, startDate.AddDate(0, 0, 1), id, 2, time.Now(), time.Now(),
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DeleteBlockById delete a room restriction
func (m *postgresDBRepo) DeleteBlockById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
	query := `delete from room_restrictions where id = $1`

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
