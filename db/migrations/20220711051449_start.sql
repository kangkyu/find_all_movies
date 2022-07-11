-- migrate:up
CREATE TABLE movies (
	id serial PRIMARY KEY,
	name text,
	cover text,
	description text
);

-- migrate:down
DROP TABLE IF EXISTS movies;
