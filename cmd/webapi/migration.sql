CREATE TABLE IF NOT EXISTS users (
	ID string PRIMARY KEY, 
	name string NOT NULL
);

CREATE TABLE IF NOT EXISTS bans (
	banisher string NOT NULL,
	banished string NOT NULL
);

CREATE TABLE IF NOT EXISTS followers (
	follower string NOT NULL,
	followed string NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
	post_ID string PRIMARY KEY,
	poster_ID string NOT NULL,
	description string,
	creation_date datetime
);

CREATE TABLE IF NOT EXISTS likes (
	post_ID string NOT NULL,
	liker string NOT NULL
);

CREATE TABLE IF NOT EXISTS comments (
	comment_ID string PRIMARY KEY,
	post_code string NOT NULL,
	user_code string NOT NULL,
	content string,
	creation_date datetime
);







