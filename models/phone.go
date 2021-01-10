package models

import (
	"database/sql"
	"time"
)

type Phone struct {
	MAC         string `db:"mac" ,json:"mac"`
	Vendor      string `db:"vendor" ,json:"vendor"`
	Model       string `db:"model" ,json:"model"`
	Description string `db:"description" ,json:"description"`

	LastSeen sql.NullTime `db:"last_seen" ,json:"last_seen"`

	ID        int `db:"id" ,json:"id"`
	CreatedAt time.Time `db:"created_at" ,json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" ,json:"updated_at"`
}

func (p *Phone) GetLines() (*[]Line, error) {
	var configuredLines []Line

	err := client.Select(&configuredLines,
		"SELECT pl.position, e.extension, e.secret, e.display_name, s.url as server " +
		"FROM phone_lines AS pl " +
		"INNER JOIN extensions AS e ON (pl.extension_id = e.id) " +
		"INNER JOIN servers AS s ON (e.server_id = s.id) " +
		"WHERE phone_id = $1 " +
		"ORDER BY position ASC;", p.ID)
	if err != nil {
		return nil, err
	}

	meta := GetPhoneMeta(p.Vendor, p.Model)

	var lines []Line
	for i := 0; i < meta.Lines; i++ {
		lineFound := false

		for _, cl := range configuredLines {
			if cl.Position == i {
				cl.Enabled = true
				lines = append(lines, cl)
				lineFound = true
				break
			}
		}

		if !lineFound {
			lines = append(lines, Line{Enabled: false})
		}
	}
	return &lines, nil
}

func GetPhone(mac, vendor, model string) (*Phone, error) {
	var phone Phone

	err := client.Get(&phone, "SELECT * FROM phones WHERE mac = $1 AND vendor = $2 AND model = $3;", mac, vendor, model)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &phone, nil
}