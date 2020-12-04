PRAGMA foreign_keys = 1;

CREATE TABLE IF NOT EXISTS user (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS tag (
    tag_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS task (
    title VARCHAR(100) NOT NULL,
    finish_dt DATETIME NOT NULL,
    tag_id INTEGER,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (tag_id) REFERENCES tag (tag_id),
    FOREIGN KEY (user_id) REFERENCES user (user_id)
);

CREATE TABLE IF NOT EXISTS current_task (
    title VARCHAR(100) NOT NULL,
    start_dt DATETIME NOT NULL,
    tag_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL UNIQUE,
    FOREIGN KEY (tag_id) REFERENCES tag (tag_id),
    FOREIGN KEY (user_id) REFERENCES user (user_id)
);