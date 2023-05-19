package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Connect opens a postgres connection. Caller must call db.Close() when done.
// A postgres driver (e.g. _ "github.com/lib/pq") must be imported in main.
func Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping: %w", err)
	}
	return db, nil
}
