package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Xivolkar/Stats/model"
	_ "github.com/denisenkom/go-mssqldb"
)

// DataStorer defines all the database operations
type DataStorer interface {
	Migrate() error
	GetStats() ([]model.Stat, error)
	Close() error
}

type dbConfig struct {
	Server       string `json:"server"`
	Port         string `json:"port"`
	UserID       string `json:"userId"`
	Password     string `json:"password"`
	DatabaseName string `json:"databaseName"`
}

// Instance - DB Instance
type Instance struct {
	DB *sql.DB
}

// Close - Shutdown DB
func (ci *Instance) Close() error {
	return ci.DB.Close()
}

func (ci *Instance) GetStats() ([]model.Stat, error) {
	var stats []model.Stat

	stmt, err := ci.DB.Prepare("select * from Stat")
	if err != nil {
		log.Println("DB Query failed")
		log.Println(err)
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error during query")
		return nil, err
	}

	defer rows.Close()

	var stat model.Stat
	for rows.Next() {
		err := rows.Scan(&stat)
		stats = append(stats, stat)
		if err != nil {
			log.Println("Error during rows scan")
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return stats, nil
}

func (ci *Instance) Migrate() error {
	return nil
}

// NewDB - Initializes the DB and returns if there is an error in the process
func NewDB(server string, port string, userid string, password string, dbName string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", server, userid, password, port, dbName)
	db, err := sql.Open("mssql", connectionString)
	return db, err
}

func LoadDbConfig() dbConfig {
	file, err := ioutil.ReadFile("./db/dbconfig.json")
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	var config dbConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}

	return config
}
