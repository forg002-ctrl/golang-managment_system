package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type DBConnectionOptions struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

type PostgresClient struct {
	db *sql.DB
}

var (
	instance *PostgresClient
	once     sync.Once
)

func (c *PostgresClient) Init(options *DBConnectionOptions) {
	once.Do(func() {
		if instance != nil {
			log.Fatal("Error: PostgresClient is already initialized")
		}

		connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			options.Host, options.Port, options.User, options.Password, options.Name)
		db, err := sql.Open("postgres", connectionString)

		if err != nil {
			log.Fatal("Error: Failed to connect to the Database: " + err.Error())
		}

		time.Sleep(2 * time.Second)
		err = db.Ping()
		if err != nil {
			log.Fatal("Error: Failed to ping the Database: " + err.Error())
		}

		c.db = db;
	})
}

func (c *PostgresClient) GetInstance() (*PostgresClient, error) {
	if instance == nil {
		return nil, fmt.Errorf("Error: PostgresClient is not initialized")
	}

	return instance, nil
}

func (c *PostgresClient) CreateTable(tableName string, tableAttributes map[string]string) error {
	if c.db == nil {
		log.Fatal("db is empty")
	}
	sqlStatement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tableName)
	
	fmt.Printf(sqlStatement)
	for columnName, columnType := range tableAttributes {
		sqlStatement += fmt.Sprintf("%s %s,", columnName, columnType)
	}

	sqlStatement = sqlStatement[:len(sqlStatement)-1] + ");"

	_, err := c.db.Exec(sqlStatement)
	if err != nil {
		return fmt.Errorf("Error: failed to create table %s: %v", tableName, err)
	}

	fmt.Printf("Table %s created successfully\n", tableName)

	return nil
}