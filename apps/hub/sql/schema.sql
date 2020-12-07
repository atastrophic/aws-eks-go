
-- name: unused
use todolist;
-- name: unused
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS todos;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS lists;
SET FOREIGN_KEY_CHECKS=1;

-- name: create-users-table
CREATE TABLE IF NOT EXISTS users (
	id CHAR(36) PRIMARY KEY,
    username VARCHAR(20) UNIQUE NOT NULL,
    password CHAR(98) NOT NULL
) CHARACTER SET utf8 COLLATE utf8_general_ci;
-- name: create-lists-table
CREATE TABLE IF NOT EXISTS lists (
	id CHAR(36) PRIMARY KEY,
    owner VARCHAR(20),
    FOREIGN KEY (owner) REFERENCES users(id)
		ON DELETE CASCADE
        ON UPDATE CASCADE
) CHARACTER SET utf8 COLLATE utf8_general_ci;
-- name: create-todos-table
CREATE TABLE IF NOT EXISTS todos (
	id CHAR(36) PRIMARY KEY,
    owner CHAR(36),
    list CHAR(36), 
    FOREIGN KEY (owner) REFERENCES users(id)
		ON DELETE SET NULL
        ON UPDATE CASCADE,
	FOREIGN KEY (list) REFERENCES lists(id)
		ON DELETE CASCADE
        ON UPDATE CASCADE,
	title VARCHAR(120),
    text VARCHAR(1000)
) CHARACTER SET utf8 COLLATE utf8_general_ci;
