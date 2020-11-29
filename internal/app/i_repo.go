package app

import (
	"time"

	"tomato/internal/models"
)

type repo interface {
	GetTasks() ([]models.TaskFromDB, error)
	SaveTask(title, tag string, date time.Time) error
	AddCurrentTask(title, tag string, date time.Time) error
	GetCurrentTask() (models.TaskFromDB, error)
	DeleteCurrentTask() error
}
