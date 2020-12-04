package repo

const (
	sqliteQueryGetTasks = `
	SELECT title, tag.name, finish_dt FROM task
	JOIN user USING (user_id)
	LEFT OUTER JOIN tag USING (tag_id)
	WHERE user.name = ?;
	`
)

const (
	sqliteQueryAddTask = `
	INSERT INTO task(title, finish_dt, tag_id, user_id)
	VALUES(?, ?, ?, ?);
	`
)

const (
	sqliteQueryGetCurrentTask = `
	SELECT title, tag.name, start_dt FROM current_task
	JOIN user USING (user_id)
	LEFT OUTER JOIN tag USING (tag_id)
	WHERE user.name = ?;
	`
)

const (
	sqliteQueryAddCurrentTask = `
	INSERT INTO current_task(title, start_dt, tag_id, user_id)
	VALUES(?, ?, ?, ?);
	`
)

const (
	sqliteQueryDeleteCurrentTask = `
	DELETE FROM current_task WHERE user_id = ?;
	`
)

const (
	sqliteQueryAddTagQuery = `
	INSERT OR IGNORE INTO tag(name) VALUES(?);
	`
)

const (
	sqliteQueryGetTagIDQuery = `
	SELECT tag_id FROM tag WHERE name = ?;
	`
)

const (
	sqliteQueryAddUserQuery = `
	INSERT OR IGNORE INTO user(name) VALUES(?);
	`
)

const (
	sqliteQueryGetUserIDQuery = `
	SELECT user_id FROM user WHERE name = ?;
	`
)
