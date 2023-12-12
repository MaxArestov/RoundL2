package repository

import (
	"fmt"
	"sync"
	"task11/pkg/models"
	"time"
)

const (
	Day   = 1
	Week  = 7
	Month = 31
)

// Repository структура, работающая с кешем, а не БД
type Repository struct {
	cache map[int][]models.Event
	mu    sync.RWMutex
}

// NewRepository создает новый репозиторий с пустым кешем
func NewRepository() *Repository {
	return &Repository{make(map[int][]models.Event), sync.RWMutex{}}
}

// CreateEvent создает новое событие и добавляет его в кеш
func (rep *Repository) CreateEvent(event models.Event) int {
	event.EventId = len(rep.cache[event.UserId]) + 1
	rep.mu.Lock()
	defer rep.mu.Unlock()
	rep.cache[event.UserId] = append(rep.cache[event.UserId], event)
	return event.EventId
}

// UpdateEvent обновляет событие в кеше
func (rep *Repository) UpdateEvent(event models.Event) error {
	for _, val := range rep.cache[event.UserId] {
		if val.EventId == event.EventId {
			rep.mu.Lock()
			val = event
			rep.mu.Unlock()
			return nil
		}

	}
	return fmt.Errorf("event %d does not exists", event.EventId)
}

// DeleteEvent удаляет событие из кеша
func (rep *Repository) DeleteEvent(userId, eventId int) error {
	for i, val := range rep.cache[userId] {
		if val.EventId == eventId {
			rep.mu.Lock()
			rep.cache[userId] = append(rep.cache[userId][:i], rep.cache[userId][i+1:]...)
			rep.mu.Unlock()
			return nil
		}
	}
	return fmt.Errorf("event %d does not exists", eventId)
}

// GetEvent получает события за указанный период и дату
func (rep *Repository) GetEvent(period, userId int, date time.Time) ([]models.Event, error) {
	if _, ok := rep.cache[userId]; !ok {
		return nil, fmt.Errorf("user with id %d does not exist", userId)
	}

	var result []models.Event
	switch period {
	case Day:
		rep.mu.Lock()
		defer rep.mu.Unlock()
		for i, val := range rep.cache[userId] {
			if val.DateTime.Year() == date.Year() &&
				val.DateTime.Month() == date.Month() &&
				val.DateTime.Day() == date.Day() {
				result = append(result, rep.cache[userId][i])
			}
		}
	case Week:
		rep.mu.Lock()
		defer rep.mu.Unlock()
		year, week := date.ISOWeek()
		for i, val := range rep.cache[userId] {
			year1, week1 := val.DateTime.ISOWeek()
			if year == year1 && week == week1 {
				result = append(result, rep.cache[userId][i])
			}
		}
	case Month:
		rep.mu.Lock()
		defer rep.mu.Unlock()
		for i, val := range rep.cache[userId] {
			if val.DateTime.Year() == date.Year() && val.DateTime.Month() == date.Month() {
				result = append(result, rep.cache[userId][i])
			}
		}
	}
	if result != nil {
		return result, nil
	} else {
		return nil, fmt.Errorf("no events for %v", date)
	}
}
