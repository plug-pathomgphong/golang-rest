package models

import (
	"database/sql"
	"time"

	"github.com/plug-pathomgphong/golang-rest/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES(?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.Id = id

	return err
}

func (e *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)
	return err
}

func (e *Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	var event Event
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	// fmt.Print(err)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *Event) Register(userId int64) error {
	query := `INSERT INTO registrations(event_id, user_id) 
	VALUES(?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)
	return err
}

func (e *Event) CancelRegister(userId int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)

	return err
}

func IsUserRegisteredForEvent(eventId, userId int64) bool {
	query := `SELECT COUNT(1) FROM registrations WHERE event_id = ? AND user_id = ?`
	row := db.DB.QueryRow(query, eventId, userId)
	var count int
	err := row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			// ไม่พบข้อมูล หมายถึงผู้ใช้ยังไม่ได้ลงทะเบียน
			return false
		}
		return false
	}
	return count > 0
}
