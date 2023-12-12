package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"task11/pkg/models"
	"task11/pkg/repository"
	"task11/pkg/service"
	"time"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Initialize() *http.ServeMux {
	mux := http.NewServeMux()

	// Обработчики по роутам
	mux.HandleFunc("/create_event", h.CreateEvent)
	mux.HandleFunc("/update_event", h.UpdateEvent)
	mux.HandleFunc("/delete_event", h.DeleteEvent)
	mux.HandleFunc("/events_for_month", h.GetMonthlyEvents)
	mux.HandleFunc("/events_for_week", h.GetWeeklyEvents)
	mux.HandleFunc("/events_for_day", h.GetDailyEvents)

	return mux
}

func (h *Handler) CreateEvent(writer http.ResponseWriter, request *http.Request) {
	event, err := ParseEvent(request) // Парсим данные события из запроса
	if err != nil {
		returnError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err = ValidateEvent(event); err != nil { // Проверяем событие на корректность
		returnError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var result []models.Event
	event.EventId = h.service.Create(event) // Создаем событие в сервисе
	result = append(result, event)
	sendJSONResponse(writer, "Successfully create event", result) // Возвращаем успешный результат
}

func (h *Handler) UpdateEvent(writer http.ResponseWriter, req *http.Request) {
	event, err := ParseEvent(req)
	if err != nil {
		returnError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err = ValidateEvent(event); err != nil {
		returnError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.service.Update(event); err != nil {
		returnError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var result []models.Event
	result = append(result, event)
	sendJSONResponse(writer, "Successfully update event", result)
}

func (h *Handler) DeleteEvent(writer http.ResponseWriter, req *http.Request) {
	event, err := ParseEvent(req)
	if err != nil {
		returnError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if event.EventId < 1 || event.UserId < 1 {
		returnError(writer, "EventId and UserId can't be 0", http.StatusBadRequest)
		return
	}

	if err = h.service.Delete(event.UserId, event.EventId); err != nil {
		returnError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	sendJSONResponse(writer, "Successfully delete event", nil)
}

func (h *Handler) GetMonthlyEvents(writer http.ResponseWriter, req *http.Request) {
	var result []models.Event

	event, err := ParseEvent(req)
	if err != nil {
		returnError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if result, err = h.service.Get(repository.Month, event.UserId, event.DateTime); err != nil {
		returnError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	sendJSONResponse(writer, "Events:\n", result)
}

func (h *Handler) GetWeeklyEvents(writer http.ResponseWriter, req *http.Request) {
	var result []models.Event

	event, err := ParseEvent(req)
	if err != nil {
		returnError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if result, err = h.service.Get(repository.Week, event.UserId, event.DateTime); err != nil {
		returnError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	sendJSONResponse(writer, "Events:\n", result)
}

func (h *Handler) GetDailyEvents(writer http.ResponseWriter, req *http.Request) {
	var result []models.Event

	event, err := ParseEvent(req)
	if err != nil {
		returnError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if result, err = h.service.Get(repository.Day, event.UserId, event.DateTime); err != nil {
		returnError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	sendJSONResponse(writer, "Events:\n", result)
}

func ParseEvent(request *http.Request) (models.Event, error) {
	var event models.Event
	var err error

	userId := request.FormValue("user_id")
	if userId != "" {
		event.UserId, err = strconv.Atoi(userId)
		if err != nil {
			return event, err
		}
	}

	event.DateTime, err = time.Parse(time.DateOnly, request.FormValue("date_time"))
	if err != nil {
		return event, err
	}

	event.Description = request.FormValue("description")

	eventId := request.FormValue("event_id")
	if eventId != "" {
		event.EventId, err = strconv.Atoi(eventId)
		if err != nil {
			return event, err
		}
	}
	return event, nil
}

func ValidateEvent(event models.Event) error {
	if event.EventId < 0 {
		return fmt.Errorf("%d: incorrect id", event.EventId)
	}

	if event.DateTime.IsZero() {
		return fmt.Errorf("%v: incorrect date", event.DateTime)
	}
	return nil
}

func returnError(writer http.ResponseWriter, error string, statusCode int) {
	ErrorResponse := struct {
		Error string `json:"error"`
	}{error}

	jsonError, err := json.Marshal(ErrorResponse)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(statusCode)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonError)
}

func sendJSONResponse(writer http.ResponseWriter, mes string, eventList []models.Event) {
	response := struct {
		Result string         `json:"result"`
		Events []models.Event `json:"events"`
	}{mes, eventList}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonData)
}
