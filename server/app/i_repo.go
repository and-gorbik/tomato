package app

import (
	"time"

	"tomato/server/models"
)

type repo interface {
	GetTasks(user string) ([]models.TaskFromDB, error)
	SaveTask(user, title, tag string, date time.Time) error
	AddCurrentTask(user, title, tag string, date time.Time) error
	GetCurrentTask(user string) (models.TaskFromDB, error)
	DeleteCurrentTask(user string) error
}
