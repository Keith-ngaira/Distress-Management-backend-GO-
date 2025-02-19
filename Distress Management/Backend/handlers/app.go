package handlers

import (
	"database/sql"
)

// App struct holds application dependencies
type App struct {
	DB *sql.DB
}
