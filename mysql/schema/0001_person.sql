-- +goose Up
CREATE TABLE person (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  first_name varchar(256) NOT NULL,
  last_name varchar(256) NOT NULL,
  middle_initial varchar(10) NULL
);

-- +goose Down
DROP TABLE person;