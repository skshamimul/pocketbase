//go:build !no_default_driver

package core

import (
	"github.com/pocketbase/dbx"
	_ "github.com/lib/pq"
)

func DefaultDBConnect(dbPath string) (*dbx.DB, error) {
	// user provided connection details:
	// server name: postgres
	// database name: postgres
	// port: 5431
	// password: MySecretPassword
	connStr := "postgresql://postgres:MySecretPassword@postgres:5431/postgres?sslmode=disable"

	db, err := dbx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
