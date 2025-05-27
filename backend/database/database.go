package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	jsonFile, err := os.Open("../utils/config.json")
	if err != nil {
		return nil, fmt.Errorf("Unable to open file: %w", err)
	}
	defer jsonFile.Close() // ensure file is closed when function exists
	data := make(map[string]interface{})

	err = json.NewDecoder(jsonFile).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	dbMap := data["database"].(map[string]interface{})
	// data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbMap["username"].(string),
		dbMap["password"].(string),
		dbMap["host"].(string),
		dbMap["port"].(string),
		dbMap["name"].(string),
	)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	return db, nil

}
