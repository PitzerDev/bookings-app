package dbrepo

import (
	"errors"
	"time"

	"github.com/tsawler/bookings-app/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if availability exists for roomID and false if no availability
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room

	return rooms, nil
}

// GetRoomByID gets a room by id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room

	if id > 2 {
		return room, errors.New("Some error")
	}

	return room, nil
}

// GetRoomByID gets a room by id
func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var user models.User

	return user, nil
}

// GetRoomByID gets a room by id
func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

// GetRoomByID gets a room by id
func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {

	return 1, "", nil
}

// AllReservations returns a slice of all reservations
func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {

	var reservations []models.Reservation

	return reservations, nil
}

// AllNewReservations returns a slice of all reservations
func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}
