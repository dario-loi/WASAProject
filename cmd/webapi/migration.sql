CREATE TABLE IF NOT EXISTS users (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  propic_id TEXT,
  bio_id TEXT,
  FOREIGN KEY(propic_id) REFERENCES photos(id),
  FOREIGN KEY(bio_id) REFERENCES bios(id),
  FOREIGN KEY(id) REFERENCES followers(followed),
  FOREIGN KEY(id) REFERENCES followers(following),
  FOREIGN KEY(id) REFERENCES bans(banisher),
  FOREIGN KEY(id) REFERENCES bans(banished),
  FOREIGN KEY(id) REFERENCES likes(liker_id)
);
CREATE TABLE IF NOT EXISTS followers (
  followed TEXT,
  following TEXT,
  PRIMARY KEY (followed, following)
);
CREATE TABLE IF NOT EXISTS bans (
  banisher TEXT,
  banished TEXT,
  PRIMARY KEY (banisher, banished)
);
CREATE TABLE IF NOT EXISTS bios (
  id TEXT PRIMARY KEY,
  bio_str TEXT,
  birth_date TEXT,
  residence TEXT,
  current_place_id TEXT,
  FOREIGN KEY(residence) REFERENCES places(id),
  FOREIGN KEY(current_place_id) REFERENCES places(id)
);
CREATE TABLE IF NOT EXISTS posts (
  id TEXT PRIMARY KEY,
  photo_id TEXT,
  description TEXT NOT NULL DEFAULT '',
  FOREIGN KEY(photo_id) REFERENCES photos(id),
  FOREIGN KEY(id) REFERENCES likes(post_id)
);
CREATE TABLE IF NOT EXISTS comment (
  id TEXT PRIMARY KEY,
  content TEXT NOT NULL DEFAULT '',
  author_id TEXT,
  creation_date TEXT NOT NULL,
  FOREIGN KEY(author_id) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS comment_list (
  comment_id TEXT,
  post_id TEXT,
  FOREIGN KEY(comment_id) REFERENCES comment(id),
  FOREIGN KEY(post_id) REFERENCES posts(id),
  PRIMARY KEY (comment_id, post_id)
);
CREATE TABLE IF NOT EXISTS occupations (
  person_id TEXT,
  name TEXT NOT NULL,
  start_date TEXT NOT NULL,
  end_date TEXT,
  place_id TEXT,
  kind TEXT CHECK( kind IN ( 'education', 'employment' ) ) NOT NULL,
  FOREIGN KEY(person_id) REFERENCES users(id),
  PRIMARY KEY (person_id, place_id) 
);
CREATE TABLE IF NOT EXISTS likes (
  liker_id TEXT,
  post_id TEXT, 
  PRIMARY KEY (liker_id, post_id) 
);
CREATE TABLE IF NOT EXISTS photos (
  id TEXT PRIMARY KEY
);
CREATE TABLE IF NOT EXISTS places (
  id TEXT PRIMARY KEY,
  country TEXT,
  state TEXT,
  city TEXT,
  lat REAL NOT NULL DEFAULT 0.0,
  long REAL NOT NULL DEFAULT 0.0,
  FOREIGN KEY(id) REFERENCES occupations(place_id)
);
