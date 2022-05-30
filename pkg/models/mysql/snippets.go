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
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Snippet{}

	// with row.Scan copy the values of the fields in sql to Snippet struct
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	}
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Will return recent snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	// Last 10 snippets
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*models.Snippet{}

	// Looping through rows like
	for rows.Next() {

		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	// To get an error happened in loop
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
