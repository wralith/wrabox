package mysql

import (
	"database/sql"

	"github.com/wralith/wrabox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into DB
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	// SQL statement that insert into snippets
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Execute statement in DB, add rest
	// Result can be '_' ignored if no need to use return value.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, nil
	}

	// Getting the ID of new inserted record
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Return specific snipped by its ID
func (m *SnippetModel) Get(id int) (*SnippetModel, error) {
	return nil, nil
}

// Will return recent snippets -10-
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
