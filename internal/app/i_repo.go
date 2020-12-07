package app

import (
	"time"

	"tomato/internal/models"
)

type repo interface {
	GetTasks(user string) ([]models.TaskResponse, error)
	SaveTask(user, title, tag string, date time.Time) error
	AddCurrentTask(user, title, tag string, date time.Time) error
	GetCurrentTask(user string) (models.TaskResponse, error)
	DeleteCurrentTask(user string) error
}
