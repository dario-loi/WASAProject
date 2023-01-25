CREATE TABLE IF NOT EXISTS users (
	ID string,
	name string
);

CREATE TABLE IF NOT EXISTS bans (
	banisher string,
	banished string
);

CREATE TABLE IF NOT EXISTS followers (
	follower string,
	followed string
);

CREATE TABLE IF NOT EXISTS posts (
	post_ID string,
	photo_code string,
	poster_ID string,
	description string,
	creation_date datetime
);

CREATE TABLE IF NOT EXISTS likes (
	post_ID string,
	liker string
);

CREATE TABLE IF NOT EXISTS comments (
	post_code string,
	user_code string,
	content string,
	creation_date datetime
);







