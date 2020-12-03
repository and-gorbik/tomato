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

func (r *Repo) GetTasks() ([]models.TaskFromDB, error) {
	rows, err := r.dbConn.Query(getTasksQuery)
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

func (r *Repo) SaveTask(title, tag string, date time.Time) error {
	return r.saveTask(saveTaskQuery, title, tag, date)
}

func (r *Repo) AddCurrentTask(title, tag string, date time.Time) error {
	return r.saveTask(addCurrentTaskQuery, title, tag, date)
}

func (r *Repo) GetCurrentTask() (models.TaskFromDB, error) {
	rows, err := r.dbConn.Query(getCurrentTaskQuery)
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

func (r *Repo) DeleteCurrentTask() error {
	if _, err := r.dbConn.Exec(deleteCurrentTaskQuery); err != nil {
		return err
	}

	return nil
}

func (r *Repo) saveTask(query, title, tag string, date time.Time) error {
	t, err := r.dbConn.Begin()
	if err != nil {
		return err
	}

	var tagID sql.NullInt32
	if err := t.QueryRow(getTagQuery, tag).Scan(&tagID); err != nil {
		if e := t.Rollback(); e != nil {
			return e
		}

		return err
	}

	if !tagID.Valid {
		if _, err = t.Exec(addTagQuery, tag); err != nil {
			if e := t.Rollback(); e != nil {
				return e
			}

			return err
		}

		if err := t.QueryRow(getTagQuery, tag).Scan(&tagID); err != nil {
			if e := t.Rollback(); e != nil {
				return e
			}

			return err
		}
	}

	if _, err = t.Exec(query, title, tag, date); err != nil {
		if e := t.Rollback(); e != nil {
			return e
		}

		return err
	}

	t.Commit()
	return nil
}
