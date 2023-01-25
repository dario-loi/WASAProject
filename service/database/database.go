/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/components"
	"github.com/sirupsen/logrus"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	SetName(name string) error
	Ping() error

	// Boilerplate code for APIs
	// each method encapsulates the logic for a specific API
	// it goes from data estracting from the DB to data serialization

	// GetUserID returns the ID of the user with the given name
	// Create the user if it doesn't exist
	PostUserID(userName string) (ID string, err error)

	// GetUserID returns the ID of the user with the given name
	// it returns an error if the user doesn't exist
	// therefore the user is NOT created.
	GetUserID(name string) (ID string, err error)

	GetUsername(ID string) (username string, err error)

	SearchUserByName(name string) (matches string, err error)

	// CheckUserExists returns true if the user with the given ID exists
	CheckUserExists(ID string) (exists bool, err error)

	// CheckPhotoExists returns true if the photo with the given ID exists
	CheckPhotoExists(ID string) (exists bool, err error)

	// CheckUsernameExists returns true if the user with the given username exists
	CheckUsernameExists(username string) (exists bool, err error)

	GetUserPhotos(ID string) (photos string, err error)

	// GetPhoto returns the profile photo of the user with the given ID
	GetUserProfile(ID string) (profile string, err error)

	GetUserFollowers(ID string) (followers string, err error)

	GetUserFollowing(ID string) (following string, err error)

	GetPhotoLikes(ID string) (likes string, err error)

	GetPhotoComments(ID string) (comments string, err error)

	GetUserBans(ID string) (bans string, err error)
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	err := db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// Load migration queries from file and execute them

	// get current directory

	cwd, err := os.Getwd()

	if err != nil {
		return nil, fmt.Errorf("error getting current directory: %w", err)
	}

	// read migration file

	migration_data, err := os.ReadFile(cwd + "/migration.sql")

	if err != nil {
		return nil, fmt.Errorf("error reading migration file: %w", err)
	}

	_, err = db.Exec(string(migration_data))

	if err != nil {
		return nil, fmt.Errorf("error executing migration: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil
}

// Wraps the Ping() method of the underlying DB connection
// This is used to check if the DB is still alive or to
// prompt the establishment of a new connection.
func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

func (db *appdbimpl) GetUsername(ID string) (username string, err error) {
	err = db.c.QueryRow(`SELECT name FROM users WHERE id = ?`, ID).Scan(&username)

	if err != nil {
		return "", fmt.Errorf("error getting username: %w", err)
	}

	return username, nil
}

// PostUserID returns the ID of the user with the given name
// Create the user if it doesn't exist

func (db *appdbimpl) PostUserID(userName string) (json string, err error) {

	// check if the user already exists
	var count int

	// Selects ALWAYS one row
	err = db.c.QueryRow(`SELECT COUNT(id) FROM users WHERE name = ?`, userName).Scan(&count)

	if err != nil {
		data, e := components.Error{Code: 500, Message: "Internal Server Error"}.ToJSON()

		if e != nil {
			return components.InternalServerError, fmt.Errorf("error converting error to JSON: %w", e)
		}

		return string(data), fmt.Errorf("error getting user ID: %w", err)
	}

	// get the count
	var userID string

	if count == 0 {
		// create the user
		// Hash the user name with SHA256

		h := sha256.New()
		h.Write([]byte(userName))
		userID = hex.EncodeToString(h.Sum(nil))

		// Insert the user in the DB
		_, err = db.c.Exec(`INSERT INTO users (id, name) VALUES (?, ?)`, userID, userName)

		if err != nil {
			data, e := components.Error{Code: 500, Message: "Internal Server Error"}.ToJSON()

			if e != nil {
				return components.InternalServerError, fmt.Errorf("error converting error to JSON: %w", e)
			}

			return string(data), fmt.Errorf("error creating user: %w", err)
		}

	} else {

		// get the user ID

		err = db.c.QueryRow(`SELECT id FROM users WHERE name = ?`, userName).Scan(&userID)

		if err != nil {
			data, e := components.Error{Code: 500, Message: "Internal Server Error"}.ToJSON()

			if e != nil {
				return components.InternalServerError, fmt.Errorf("error converting error to JSON: %w", e)
			}

			return string(data), fmt.Errorf("error getting user ID: %w", err)
		}

	}

	// return the user ID
	userID_json := components.SHA256hash{Hash: userID}

	data, err := userID_json.ToJSON()

	if err != nil {
		data, e := components.Error{Code: 500, Message: "Internal Server Error"}.ToJSON()

		if e != nil {
			return components.InternalServerError, fmt.Errorf("error converting error to JSON: %w", e)
		}

		return string(data), fmt.Errorf("error converting user to JSON: %w", err)
	}

	return string(data), nil

}

func (db *appdbimpl) SearchUserByName(name string) (matches string, err error) {

	res, err := db.c.Query(`SELECT u.name FROM users as u WHERE u.name LIKE '%'||?||'%'`, name)

	defer func() {
		if res != nil {
			err := res.Close()
			if err != nil {
				logrus.Errorf("error closing result set: %v", err)
			}
		}
	}()

	if err != nil {
		return components.InternalServerError, fmt.Errorf("error searching user: %w", err)
	}

	var users []components.User

	for res.Next() {

		user := components.User{}

		err = res.Scan(&user.Uname)

		if err != nil {
			return components.InternalServerError, fmt.Errorf("error scanning user: %w", err)
		}

		users = append(users, user)
	}

	data, err := json.Marshal(users)

	if err != nil {
		return components.InternalServerError, fmt.Errorf("error converting users to JSON: %w", err)
	}

	return string(data), nil

}

func (db *appdbimpl) CheckUserExists(userID string) (exists bool, err error) {

	var count int

	// Selects ALWAYS one row
	err = db.c.QueryRow(`SELECT COUNT(id) FROM users WHERE id = ?`, userID).Scan(&count)

	if err != nil {
		return false, fmt.Errorf("error getting user ID: %w", err)
	}

	if count == 0 {
		return false, fmt.Errorf("user %s does not exist", userID)
	}

	return true, nil

}

func (db *appdbimpl) CheckPhotoExists(photoID string) (exists bool, err error) {

	var count int

	// Selects ALWAYS one row
	err = db.c.QueryRow(`SELECT COUNT(photo_code) FROM posts WHERE photo_code = ?`, photoID).Scan(&count)

	if err != nil {
		return false, fmt.Errorf("error getting photo ID: %w", err)
	}

	if count == 0 {
		return false, fmt.Errorf("photo %s does not exist", photoID)
	}

	return true, nil

}

func (db *appdbimpl) CheckUsernameExists(username string) (exists bool, err error) {

	var count int

	// Selects ALWAYS one row
	err = db.c.QueryRow(`SELECT COUNT(ID) FROM users WHERE name = ?`, username).Scan(&count)

	if err != nil {
		return false, fmt.Errorf("error getting user ID: %w", err)
	}

	if count == 0 {
		return false, fmt.Errorf("user %s does not exist", username)
	}

	return true, nil

}

func (db *appdbimpl) GetUserPhotos(userID string) (photo string, err error) {

	photoIDlist := components.IDList{}

	res, err := db.c.Query(`SELECT ps.photo_name FROM posts AS pt 
		WHERE ps.poster_name = ?`, userID)

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error getting user's photos: %w", err)
	}

	for res.Next() {
		var photoID components.SHA256hash

		err = res.Scan(&photoID.Hash)

		if err != nil {
			return components.InternalServerError,
				fmt.Errorf("error scanning photo: %w", err)
		}

		photoIDlist.IDs = append(photoIDlist.IDs, photoID)
	}

	data, err := photoIDlist.ToJSON()

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error converting photo to JSON: %w", err)
	}

	return string(data), nil
}

func (db *appdbimpl) GetUserID(name string) (ID string, err error) {

	var userID string

	err = db.c.QueryRow(`SELECT id FROM users WHERE name = ?`, name).Scan(&userID)

	if err != nil {
		return components.InternalServerError, fmt.Errorf("error getting user ID: %w", err)
	}

	return userID, nil

}

func (db *appdbimpl) GetUserProfile(userID string) (profile string, err error) {

	var username string

	err = db.c.QueryRow(`SELECT name FROM Users WHERE id = ?`, userID).Scan(&username)

	if err != nil {
		return components.InternalServerError, fmt.Errorf("error getting username: %w", err)
	}

	photoIDlist := []components.SHA256hash{}

	res, err := db.c.Query(`SELECT ps.photo_name FROM posts AS pt 
		WHERE ps.poster_name = ?`, userID)

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error getting user's photos: %w", err)
	}

	for res.Next() {
		var photoID components.SHA256hash

		err = res.Scan(&photoID.Hash)

		if err != nil {
			return components.InternalServerError,
				fmt.Errorf("error scanning photo: %w", err)
		}

		photoIDlist = append(photoIDlist, photoID)
	}

	prof_struct := components.Profile{
		Username: username,
		Photos:   photoIDlist,
	}

	data, err := prof_struct.ToJSON()

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error converting profile to JSON: %w", err)
	}

	return string(data), nil

}

func (db *appdbimpl) GetUserFollowers(userID string) (followers string, err error) {

	res, err := db.c.Query(`SELECT follower FROM followers WHERE followed = ?`, userID)

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error getting user's followers: %w", err)
	}

	var followerNames []string

	for res.Next() {
		var followerID string

		err = res.Scan(&followerID)

		if err != nil {
			return components.InternalServerError,
				fmt.Errorf("error scanning follower: %w", err)
		}

		followerName, err := db.GetUsername(followerID)

		if err != nil {
			return components.InternalServerError, fmt.Errorf("error getting follower name: %w", err)
		}

		followerNames = append(followerNames, followerName)

	}

	data, err := json.MarshalIndent(followerNames, "", "	")

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error converting followers to JSON: %w", err)
	}

	return string(data), nil
}

func (db *appdbimpl) GetUserFollowing(userID string) (following string, err error) {

	res, err := db.c.Query(`SELECT followed FROM followers WHERE follower = ?`, userID)

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error getting user's following: %w", err)
	}

	var followingNames []string

	for res.Next() {
		var followingID string

		err = res.Scan(&followingID)

		if err != nil {
			return components.InternalServerError,
				fmt.Errorf("error scanning following: %w", err)
		}

		followingName, err := db.GetUsername(followingID)

		if err != nil {
			return components.InternalServerError, fmt.Errorf("error getting following name: %w", err)
		}

		followingNames = append(followingNames, followingName)

	}

	data, err := json.MarshalIndent(followingNames, "", "	")

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error converting following to JSON: %w", err)
	}

	return string(data), nil
}

func (db *appdbimpl) GetPhotoLikes(photoID string) (likes string, err error) {

	res, err := db.c.Query(`SELECT l.liker FROM likes as l, posts as p WHERE p.photo_code = ? AND p.post_ID = l.likes`, photoID)

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error getting photo's likes: %w", err)
	}

	var likerNames []string

	for res.Next() {
		var likerID string

		err = res.Scan(&likerID)

		if err != nil {
			return components.InternalServerError,
				fmt.Errorf("error scanning liker: %w", err)
		}

		likerName, err := db.GetUsername(likerID)

		if err != nil {
			return components.InternalServerError, fmt.Errorf("error getting liker name: %w", err)
		}

		likerNames = append(likerNames, likerName)

	}

	data, err := json.MarshalIndent(likerNames, "", "	")

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error converting likes to JSON: %w", err)
	}

	return string(data), nil
}

func (db *appdbimpl) GetPhotoComments(photoID string) (comments string, err error) {

	res, err := db.c.Query(`SELECT u.name, c.content, c.creation_date, c.post_code FROM comments as c, posts as p, users as u WHERE p.photo_code = ? AND p.post_ID = c.post_code AND u.ID = p.poster_ID`, photoID)

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error getting photo's comments: %w", err)
	}

	var commentsList []components.Comment

	for res.Next() {

		var comment components.Comment

		err = res.Scan(&comment.Username, &comment.Body, &comment.CreationTime, &comment.Parent)

		if err != nil {
			return components.InternalServerError,
				fmt.Errorf("error scanning comment: %w", err)
		}

	}

	data, err := json.MarshalIndent(commentsList, "", "	")

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error converting comments to JSON: %w", err)
	}

	return string(data), nil
}

func (db *appdbimpl) GetUserBans(username string) (bans string, err error) {

	res, err := db.c.Query(`SELECT b.banished FROM bans as b, users as u WHERE u.name = ? AND u.ID = b.banisher`, username)

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error getting user's bans: %w", err)
	}

	var banList []components.User

	for res.Next() {

		var ban components.User

		err = res.Scan(&ban.Uname)

		if err != nil {
			return components.InternalServerError,
				fmt.Errorf("error scanning ban record: %w", err)
		}

	}

	data, err := json.MarshalIndent(banList, "", "	")

	if err != nil {
		return components.InternalServerError,
			fmt.Errorf("error converting bans to JSON: %w", err)
	}

	return string(data), nil
}
