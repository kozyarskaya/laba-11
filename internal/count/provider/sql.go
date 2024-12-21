package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) SelectCount() (string, error) {
	var msg string
	err := p.conn.QueryRow("SELECT summa FROM count ORDER BY id DESC LIMIT 1").Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return msg, nil
}

func (p *Provider) InsertCount(v float32) error {
	_, err := p.conn.Exec("INSERT INTO count (val, summa) VALUES ($1, $1 + (SELECT COALESCE(summa, 0) FROM count ORDER BY id DESC LIMIT 1))", v)
	if err != nil {
		return err
	}

	return nil
}

func (p *Provider) UpdateCount(v float32) error {
	_, err := p.conn.Exec("UPDATE count SET val = $1, summa = (val + (SELECT summa FROM count WHERE id = ((SELECT MAX(id) FROM count) - 1))) WHERE id = (SELECT MAX(id) FROM count)", v)
	if err != nil {
		return err
	}
	return nil
}
