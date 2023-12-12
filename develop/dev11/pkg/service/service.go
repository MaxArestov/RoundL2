package service

import (
	"task11/pkg/models"
	"task11/pkg/repository"
	"time"
)

type Service struct {
	rep *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{rep: r}
}

func (s *Service) Create(e models.Event) int {
	return s.rep.CreateEvent(e)
}

func (s *Service) Update(e models.Event) error {
	return s.rep.UpdateEvent(e)
}

func (s *Service) Get(period, userId int, date time.Time) ([]models.Event, error) {
	return s.rep.GetEvent(period, userId, date)
}

func (s *Service) Delete(userId, eventId int) error {
	return s.rep.DeleteEvent(userId, eventId)
}
