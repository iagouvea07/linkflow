CREATE DATABASE IF NOT EXISTS linkflow;

USE linkflow;

CREATE TABLE IF NOT EXISTS position (current_position int);

INSERT INTO position (current_position)
SELECT 100000
WHERE NOT EXISTS (SELECT 1 FROM position);

CREATE TABLE IF NOT EXISTS addresses (
    id INT NOT NULL PRIMARY KEY,
    url VARCHAR(60),
    encoded VARCHAR(7)
);