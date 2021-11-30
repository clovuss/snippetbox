package postgresql

import (
	"context"
	"errors"
	"github.com/clovuss/snippetbox/pkg/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (s *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `select * from snippets where id=$1`
	row := s.DB.QueryRow(context.Background(), stmt, id)
	r := &models.Snippet{}
	err := row.Scan(&r.Id, &r.Title, &r.Content)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return r, nil

}

func (s *SnippetModel) Insert(title string, content string) error {
	stmt := `INSERT INTO snippets (title, content)
    VALUES($1, $2)`
	_, err := s.DB.Exec(context.Background(), stmt, title, content)
	if err != nil {
		return err
	}
	return nil
}

func (s *SnippetModel) Getlatest() ([]*models.Snippet, error) {
	stmt := ` select * from snippets order by id desc limit 3 `
	rows, err := s.DB.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	snippets := make([]*models.Snippet, 0)
	for rows.Next() {
		s := &models.Snippet{}
		err := rows.Scan(&s.Id, &s.Title, &s.Content)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil

}
