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
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	SetName(name string) error
	// QueryAndPack executes the query and pakcs the result into the provided interface.
	// This allows for a *very* generic way of querying the DB and getting back the result.
	//
	// This proper use of polymorphism will probably be branded as "bad practice" by the review
	// and be promptly converted to a hundred different (but equal) methods, with complimentary
	// `switch` statements, just like the good old days!
	QueryAndPack(data_struct *interface{}, query string, args ...interface{}) ([]interface{}, error)

	// ExecQuery executes the query and returns an error if the query failed or no rows were affected.
	// Another generic method to allow for execution of queries that do *not* have a return value.
	ExecQuery(query string, args ...interface{}) error
	Ping() error
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

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
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

func (db *appdbimpl) ExecQuery(query string, args ...interface{}) error {
	res, err := db.c.Exec(query, args...)

	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()

	if err != nil {
		return err
	} else if affected == 0 {
		return errors.New("no rows affected") // convert to custom error to be properly handled into http response
	}

	return nil
}

func (db *appdbimpl) QueryAndPack(data_struct *interface{}, query string, args ...interface{}) ([]interface{}, error) {

	rows, err := db.c.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	cols, err := rows.ColumnTypes()

	if err != nil {
		return nil, err
	}

	reflected := reflect.ValueOf(data_struct).Elem()

	for idx, col := range cols {
		if reflected.FieldByIndex([]int{idx}).Type() != col.ScanType() {
			return nil, errors.New("type mismatch")
		}
	}

	//create slice of structs of the same type as the provided struct
	//100% safer than generics
	ret_struct := reflect.MakeSlice(reflect.SliceOf(reflected.Type()), 0, len(cols))

	for rows.Next() {
		//101% safer than generics
		row_struct := reflect.New(reflected.Type()).Interface()
		err = rows.Scan(row_struct)
		if err != nil {
			return nil, err
		}

		ret_struct = reflect.Append(ret_struct, reflect.ValueOf(row_struct).Elem())
	}

	//102% safer than generics
	return ret_struct.Interface().([]interface{}), nil

}
