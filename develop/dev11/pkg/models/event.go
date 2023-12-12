package models

import "time"

// Event структура для парсинга в http-req-res.
type Event struct {
	EventId     int       `json:"event_id"`
	UserId      int       `json:"user_id"`
	DateTime    time.Time `json:"date_time"`
	Description string    `json:"description"`
}
