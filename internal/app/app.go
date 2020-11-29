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
	repo               repo
}

func New(settings *models.Settings, repo repo) *App {
	return &App{
		editorExecutable:   settings.EditorPath,
		currentTaskFile:    settings.CurrentTasksPath,
		smallBreakDuration: time.Duration(settings.Tomato.SmallBreakMinutes) * time.Minute,
		bigBreakDuration:   time.Duration(settings.Tomato.BigBreakMinutes) * time.Minute,
		workDuration:       time.Duration(settings.Tomato.WorkingMinutes) * time.Minute,
		repo:               repo,
	}
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
