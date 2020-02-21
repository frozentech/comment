package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

// Connection ...
var Connection *pgx.Conn

// Connect ...
func Connect() (conn *pgx.Conn, err error) {
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("unable to connection to database: %v", err)
	}

	return
}
