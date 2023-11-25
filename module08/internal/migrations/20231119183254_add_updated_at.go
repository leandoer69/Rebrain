package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddUpdatedAt, downAddUpdatedAt)
}

func upAddUpdatedAt(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE users ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT now();")
	if err != nil {
		return err
	}
	return nil
}

func downAddUpdatedAt(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("ALTER TALBE users DROP COLUMN IF EXISTS updated_at;")
	if err != nil {
		return err
	}
	return nil
}
