package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) SelectName() (string, error) {
	var msg string
	err := p.conn.QueryRow("SELECT name FROM query ORDER BY id DESC LIMIT 1").Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return msg, nil
}

func (p *Provider) InsertQuery(msg string) error {
	_, err := p.conn.Exec("INSERT INTO query (name) VALUES ($1)", msg)
	if err != nil {
		return err
	}

	return nil
}

func (p *Provider) UpdateQuery(n string) error {
	_, err := p.conn.Exec("UPDATE query SET name = $1 WHERE id = (SELECT MAX(id) FROM query)", n)
	if err != nil {
		return err
	}
	return nil
}
