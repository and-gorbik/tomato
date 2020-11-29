package repo

const (
	getTasksQuery          = `SELECT title, tag.name, finish_date FROM task INNER JOIN tag USING tag_id`
	saveTaskQuery          = `INSERT INTO task(title, tag, finish_date) VALUES (?, ?, ?);`
	getCurrentTaskQuery    = `SELECT title, tag.name, start_date FROM current_task INNER JOIN tag USING tag_id;`
	addCurrentTaskQuery    = `INSERT INTO current_task(title, tag_id, start_date) VALUES (?, ?, ?);`
	deleteCurrentTaskQuery = `DELETE FROM current_task;`
	addTagQuery            = `INSERT INTO tag(name) VALUES (?);`
	getTagQuery            = `SELECT tag_id FROM tag WHERE name = ?;`
)