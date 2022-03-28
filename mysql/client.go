package mysql

import (
	"context"
	"os"

	"github.com/jmoiron/sqlx"
)

// Client holds a reference to a mysql DB.
type Client struct {
	db *sqlx.DB
}

// Connect attempts a connection the mysql db defined in the MYSQL_CONNECTION_URI environment variable.
func Connect(ctx context.Context) (*Client, error) {
	db, err := sqlx.ConnectContext(ctx, "mysql", os.Getenv("MYSQL_CONNECTION_URI"))
	if err != nil {
		return nil, err
	}

	return &Client{
		db: db,
	}, nil
}
