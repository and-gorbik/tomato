package app

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"tomato/internal/infrastructure"
	"tomato/internal/models"
)

type App struct {
	repo              repo
	taskfile          TaskFile
	user              string
	defaultTag        string
	editorPath        string
	smallBreakMinutes int
	bigBreakMinutes   int
	workingMinutes    int
}

func New(settings *models.Settings, repo repo, user string) *App {
	return &App{
		repo:              repo,
		user:              user,
		taskfile:          TaskFile{Path: settings.CurrentTasksPath},
		defaultTag:        settings.DefaultTag,
		editorPath:        settings.EditorPath,
		smallBreakMinutes: settings.Tomato.SmallBreakMinutes,
		bigBreakMinutes:   settings.Tomato.BigBreakMinutes,
		workingMinutes:    settings.Tomato.WorkingMinutes,
	}
}

func (a *App) LoadTasks(fname string, forced bool) (err error) {
	return a.taskfile.Load(fname, forced)
}

func (a *App) EditTasks() error {
	cmd := exec.Command(a.editorPath, a.taskfile.Path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func (a *App) StartTask() error {
	t, err := a.taskfile.Next()
	if err != nil {
		return err
	}

	if err = a.repo.AddCurrentTask(a.user, t.Title, t.Tag, time.Now()); err != nil {
		return err
	}

	return infrastructure.PlanNotifies()
}

func (a *App) StopTask() error {
	t, err := a.repo.GetCurrentTask(a.user)
	if err != nil {
		return err
	}

	if err = a.taskfile.Prepend(t.Title, t.Tag); err != nil {
		return err
	}

	if err := a.repo.DeleteCurrentTask(a.user); err != nil {
		return err
	}

	return infrastructure.RemoveNotifies()
}

func (a *App) GetLog() error {
	tasks, err := a.repo.GetTasks(a.user)
	if err != nil {
		return err
	}

	for _, t := range tasks {
		tag := ""
		if t.Tag != nil {
			tag = *t.Tag
		}
		fmt.Printf("%s\t\t%s\t\t%v\n", t.Title, tag, t.Date)
	}

	return nil
}

func (a *App) GetCurrentTasks() error {
	return a.taskfile.Print()
}
