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
	"os"
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
	QueryAndPack(data_struct interface{}, query string, args ...interface{}) ([]interface{}, error)

	// ExecQuery executes the query and returns an error if the query failed or no rows were affected.
	// Another generic method to allow for execution of queries that do *not* have a return value.
	ExecQuery(query string, args ...interface{}) error

	QueryAndPackRow(data_struct interface{}, query string, args ...interface{}) error
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

func (db *appdbimpl) QueryAndPack(data_struct interface{}, query string, args ...interface{}) ([]interface{}, error) {

	// Get the type of the data_struct
	data_struct_type := reflect.TypeOf(data_struct).Elem()

	// Get the number of fields in the data_struct
	data_struct_num_fields := data_struct_type.NumField()

	// Get the number of arguments
	num_args := len(args)

	// Create a slice of interfaces to hold the arguments
	// This is needed because the `Query` method expects a slice of interfaces
	// as arguments.
	// The arguments are passed as a variadic parameters, so we need to convert
	// them to a slice of interfaces.
	args_slice := make([]interface{}, num_args)
	for i := 0; i < num_args; i++ {
		args_slice[i] = args[i]
	}

	// Execute the query
	rows, err := db.c.Query(query, args_slice...)

	if err != nil {
		return nil, err
	}

	// Create a slice of interfaces to hold the result
	var result []interface{}

	// Iterate over the rows
	for rows.Next() {

		// Create a slice of interfaces to hold the fields
		fields := make([]interface{}, data_struct_num_fields)

		// Iterate over the fields
		for i := 0; i < data_struct_num_fields; i++ {
			// Get the field type
			field_type := data_struct_type.Field(i).Type

			// Create a new pointer to the field type
			field := reflect.New(field_type)

			// Set the field in the slice of interfaces
			fields[i] = field.Interface()
		}

		// Scan the row into the fields
		err = rows.Scan(fields...)

		if err != nil {
			return nil, err
		}

		// Create a new instance of the data_struct
		new_data_struct := reflect.New(data_struct_type)

		// Iterate over the fields
		for i := 0; i < data_struct_num_fields; i++ {
			// Get the field type
			field_type := data_struct_type.Field(i).Type

			// Get the field value
			field_value := reflect.ValueOf(fields[i]).Elem()

			// Set the field in the new instance of the data_struct
			new_data_struct.Elem().Field(i).Set(field_value.Convert(field_type))
		}

		// Append the new instance of the data_struct to the result
		result = append(result, new_data_struct.Interface())
	}

	return result, nil
}

func (db *appdbimpl) QueryAndPackRow(data_struct interface{}, query string, args ...interface{}) error {

	// Get the type of the data_struct
	data_struct_type := reflect.TypeOf(data_struct).Elem()

	// Get the number of fields in the data_struct
	data_struct_num_fields := data_struct_type.NumField()

	// Get the number of arguments
	num_args := len(args)

	// Create a slice of interfaces to hold the arguments
	// This is needed because the `Query` method expects a slice of interfaces
	// as arguments.
	// The arguments are passed as a variadic parameters, so we need to convert
	// them to a slice of interfaces.
	args_slice := make([]interface{}, num_args)
	for i := 0; i < num_args; i++ {
		args_slice[i] = args[i]
	}

	// Execute the query
	rows, err := db.c.Query(query, args_slice...)

	if err != nil {

		return err
	}

	// Iterate over the rows
	for rows.Next() {

		// Create a slice of interfaces to hold the fields
		fields := make([]interface{}, data_struct_num_fields)

		// Iterate over the fields
		for i := 0; i < data_struct_num_fields; i++ {
			// Get the field type
			field_type := data_struct_type.Field(i).Type

			// Create a new pointer to the field type
			field := reflect.New(field_type)

			// Set the field in the slice of interfaces
			fields[i] = field.Interface()
		}

		// Scan the row into the fields
		err = rows.Scan(fields...)

		if err != nil {
			return err
		}

		// Create a new instance of the data_struct
		new_data_struct := reflect.New(data_struct_type)

		// Iterate over the fields
		for i := 0; i < data_struct_num_fields; i++ {
			// Get the field type
			field_type := data_struct_type.Field(i).Type

			// Get the field value
			field_value := reflect.ValueOf(fields[i]).Elem()

			// Set the field in the new instance of the data_struct
			new_data_struct.Elem().Field(i).Set(field_value.Convert(field_type))
		}

		// Set the new instance of the data_struct to the data_struct
		data_struct = new_data_struct.Interface()

		return nil
	}

	return errors.New("no rows found")

}
