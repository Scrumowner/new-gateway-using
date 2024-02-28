package migrator

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"user_service/internal/models"
)

type Migrator struct {
	db *sqlx.DB
}

func NewMigrator(db *sqlx.DB) *Migrator {
	return &Migrator{
		db: db,
	}
}
func (m *Migrator) Migrate(tabler models.User) {
	tn := tabler.TableName()
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ( username text, email text, password text);", tn)
	m.db.Exec(query)
}
