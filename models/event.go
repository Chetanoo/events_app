package models

import (
	"events_app/db"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events (name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	resultId, err := result.LastInsertId()
	e.Id = resultId
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
		SELECT * FROM events	
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
		SELECT * FROM events WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e *Event) Delete() error {
	query := `
		DELETE FROM events WHERE id = ?
	`
	_, err := db.DB.Exec(query, e.Id)
	return err
}

func (e *Event) Update() error {
	query := `
		UPDATE events SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ? WHERE id = ?
	`
	_, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.Id)

	return err
}

func (e *Event) Register(userId int64) error {
	query := `
		INSERT INTO registrations (event_id, user_id) values (?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.Id, userId)

	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := `
		DELETE FROM registrations WHERE event_id = ? AND user_id = ?
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.Id, userId)
	return err
}
