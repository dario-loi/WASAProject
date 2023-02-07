
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS users (
	ID string PRIMARY KEY NOT NULL, 
	name string UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS bans (
	banisher string NOT NULL,
	banished string NOT NULL,
	FOREIGN KEY (banisher) REFERENCES users(ID) ON DELETE CASCADE ON UPDATE CASCADE,
	FOREIGN KEY (banished) REFERENCES users(ID) ON DELETE CASCADE ON UPDATE CASCADE

);

CREATE TABLE IF NOT EXISTS followers (
	follower string NOT NULL,
	followed string NOT NULL,
	FOREIGN KEY (follower) REFERENCES users(ID) ON DELETE CASCADE ON UPDATE CASCADE,
	FOREIGN KEY (followed) REFERENCES users(ID) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS posts (
	post_ID string PRIMARY KEY,
	poster_ID string NOT NULL,
	description string,
	creation_date datetime,
	FOREIGN KEY (poster_ID) REFERENCES users(ID) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS likes (
	post_ID string NOT NULL,
	liker string NOT NULL,
	FOREIGN KEY (post_ID) REFERENCES posts(post_ID) ON DELETE CASCADE ON UPDATE CASCADE,
	FOREIGN KEY (liker) REFERENCES users(ID) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS comments (
	comment_ID string PRIMARY KEY,
	post_code string NOT NULL,
	user_code string NOT NULL,
	content string,
	creation_date datetime,
	FOREIGN KEY (post_code) REFERENCES posts(post_ID) ON DELETE CASCADE ON UPDATE CASCADE,
	FOREIGN KEY (user_code) REFERENCES users(ID) ON DELETE CASCADE ON UPDATE CASCADE
);







