/* CREATE TABLES */
CREATE TABLE IF NOT EXISTS projects (
    id            INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
    name          VARCHAR(80) NOT NULL,
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tasks (
    id            INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
    project_id    INTEGER,		
    title         VARCHAR(80) NOT NULL,
    description   VARCHAR(512) NOT NULL,
    status_id     INTEGER,
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS notes (
    id            INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
    task_id       INTEGER NOT NULL,
    title         VARCHAR(80) NOT NULL,
    description   VARCHAR(200) NOT NULL,
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS status (
    id            INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
    title         VARCHAR(30) NOT NULL,
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS logs (
    id            INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
    code          VARCHAR(30) NOT NULL,   /* codes for better grouping */
    action        VARCHAR(30) NOT NULL,   /* created, modified, deleted, ... */
    caller        VARCHAR(100) NOT NULL,  /* caller is a combination of "{project|task}:{project|task}.id" */
    message       VARCHAR(512) NOT NULL,
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO logs ("code", "action", "caller", "message") VALUES ("MIGRATION", "created", "SYSTEM", "Ran migration: 00_init_tables.sql");
