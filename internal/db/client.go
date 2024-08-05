package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"time"
)

// DBClient represents a client to the database with a connection status.
type DBClient struct {
	DB              *sql.DB
	Connected       bool
	ConnectionError error
}

// NewDBClient initializes a new database client.
func NewDBClient() DBClient {
	FirstAttemptTime := time.Now()
	client := DBClient{}
	err := client.Connect()
	for err != nil {
		duration := time.Now().Sub(FirstAttemptTime)
		fmt.Printf("[%s] Failed to connect to db.. Retrying...\n", fmtDuration(duration))
		err = client.Connect()
	}
	client.Connected = true
	return client
}

// Connect establishes a connection to the database.
func (client *DBClient) Connect() error {
	db, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		client.ConnectionError = err
		return err
	}

	// Try to make a connection
	err = db.Ping()
	if err != nil {
		client.ConnectionError = err
		return err
	}

	client.DB = db
	client.Connected = true
	client.ConnectionError = nil
	return nil
}

// Close terminates the connection to the database.
func (client *DBClient) Close() error {
	if client.DB != nil {
		return client.DB.Close()
	}
	return nil
}

func getConnectionString() string {
	user := os.Getenv("lowLatencyWebUser")
	password := os.Getenv("lowLatencyWebPassword")
	host := os.Getenv("lowLatencyWebHost")
	port := os.Getenv("lowLatencyWebPort")
	dbname := os.Getenv("lowLatencyWebDatabase")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}
