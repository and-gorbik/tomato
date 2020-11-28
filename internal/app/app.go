package app

import (
	"time"
	"tomato/internal/models"
)

type App struct {
	editorExecutable   string
	currentTaskFile    string
	smallBreakDuration time.Duration
	bigBreakDuration   time.Duration
	workDuration       time.Duration
}

func New(cfg *models.Config) *App {
	return &App{}
}

func (a *App) LoadTasks(filename string, forced bool) (err error) {
	return
}

func (a *App) EditTasks() (err error) {
	return
}

func (a *App) StartTask(task string) (err error) {
	return
}

func (a *App) StopTask() (err error) {
	return
}

func (a *App) GetLog() (err error) {
	return
}

func (a *App) GetCurrentTasks() (err error) {
	return
}
