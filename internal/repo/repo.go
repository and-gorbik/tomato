package repo

import (
	"database/sql"
	"time"

	"tomato/server/models"
)

type Repo struct {
	dbConn *sql.DB
}

func New(dbconn *sql.DB) *Repo {
	return &Repo{
		dbConn: dbconn,
	}
}

func (r *Repo) GetTasks(user string) ([]models.TaskFromDB, error) {
	rows, err := r.dbConn.Query(sqliteQueryGetTasks, user)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]models.TaskFromDB, 0)
	for rows.Next() {
		var t models.TaskFromDB
		if err := rows.Scan(&t.Title, &t.Tag, &t.Date); err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Repo) SaveTask(user, title, tag string, date time.Time) error {
	return r.saveTask(sqliteQueryAddTask, title, tag, user, date)
}

func (r *Repo) AddCurrentTask(user, title, tag string, date time.Time) error {
	return r.saveTask(sqliteQueryAddCurrentTask, title, tag, user, date)
}

func (r *Repo) GetCurrentTask(user string) (models.TaskFromDB, error) {
	rows, err := r.dbConn.Query(sqliteQueryGetCurrentTask, user)
	if err != nil {
		return models.TaskFromDB{}, err
	}

	defer rows.Close()
	if !rows.Next() {
		return models.TaskFromDB{}, nil
	}

	var task models.TaskFromDB
	if err := rows.Scan(&task.Title, &task.Tag, &task.Date); err != nil {
		return models.TaskFromDB{}, nil
	}

	return task, nil
}

func (r *Repo) DeleteCurrentTask(user string) error {
	if _, err := r.dbConn.Exec(sqliteQueryDeleteCurrentTask, user); err != nil {
		return err
	}

	return nil
}

func (r *Repo) saveTask(query, title, tag, user string, date time.Time) error {
	t, err := r.dbConn.Begin()
	if err != nil {
		return err
	}

	rollback := func(err error) error {
		if e := t.Rollback(); e != nil {
			return e
		}

		return err
	}

	if _, err = t.Exec(sqliteQueryAddTagQuery, tag); err != nil {
		return rollback(err)
	}

	if _, err = t.Exec(sqliteQueryAddUserQuery, user); err != nil {
		return rollback(err)
	}

	var tagID sql.NullInt32
	if err = t.QueryRow(sqliteQueryGetTagIDQuery, tag).Scan(&tagID); err != nil {
		return rollback(err)
	}

	var userID int32
	if err = t.QueryRow(sqliteQueryGetUserIDQuery, user).Scan(&userID); err != nil {
		return rollback(err)
	}

	if _, err = t.Exec(query, title, date, tagID, userID); err != nil {
		return rollback(err)
	}

	t.Commit()
	return nil
}
