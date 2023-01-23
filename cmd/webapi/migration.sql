CREATE TABLE IF NOT EXISTS Users (
	ID string,
	name string
);

CREATE TABLE IF NOT EXISTS Bans (
	banisher string,
	banished string
);

CREATE TABLE IF NOT EXISTS Follower (
	follower string,
	followed string
);

CREATE TABLE IF NOT EXISTS Posts (
	post_ID string,
	photo_code string,
	poster_ID string,
	description string,
	creation_date datetime
);

CREATE TABLE IF NOT EXISTS Likes (
	post_ID string,
	liker string
);

CREATE TABLE IF NOT EXISTS Comments (
	post_code string,
	user_code string,
	content string,
	creation_date datetime
);







