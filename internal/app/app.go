package app

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"tomato/internal/models"
)

type App struct {
	repo              repo
	defaultTag        string
	editorPath        string
	currentTasksPath  string
	smallBreakMinutes int
	bigBreakMinutes   int
	workingMinutes    int
}

func New(settings *models.Settings, repo repo) *App {
	return &App{
		repo:              repo,
		defaultTag:        settings.DefaultTag,
		editorPath:        settings.EditorPath,
		currentTasksPath:  settings.CurrentTasksPath,
		smallBreakMinutes: settings.Tomato.SmallBreakMinutes,
		bigBreakMinutes:   settings.Tomato.BigBreakMinutes,
		workingMinutes:    settings.Tomato.WorkingMinutes,
	}
}

func (a *App) LoadTasks(fname string, forced bool) (err error) {
	src, err := os.Open(fname)
	if err != nil {
		return
	}

	if _, err = csv.NewReader(src).ReadAll(); err != nil {
		return err
	}

	if _, err = src.Seek(0, 0); err != nil {
		return err
	}

	defer src.Close()

	dst, err := os.OpenFile(a.currentTasksPath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return
	}

	defer dst.Close()

	dstInfo, err := dst.Stat()
	if err != nil {
		return
	}

	if dstInfo.Size() != 0 && !forced {
		err = fmt.Errorf("Current tasks is not finished")
		return
	}

	_, err = io.Copy(dst, src)
	return
}

func (a *App) EditTasks() (err error) {
	return
}

func (a *App) StartTask() (err error) {
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
