package statemanager

import (
	"database/sql"
	"encoding/json"
	"github.com/psolru/terrastate-http/sqlite3"
)

type TfState struct {
	Ident   string `db:"ident"`
	Version uint32 `json:"Version" db:"version"`
	Serial  uint32 `json:"Serial" db:"serial"`
	Raw     string `db:"data"`
	Lock    uint8  `db:"lock"`
	LockID  string `db:"lock_id"`
}

// Get return the terraform state specified by ident
func (tfState *TfState) Get() error {
	row := sqlite3.QueryRow("SELECT `ident`, `version`, `serial`, `data`, `lock`, `lock_id` FROM tf_state WHERE ident = $1", tfState.Ident)
	err := row.Scan(&tfState.Ident, &tfState.Version, &tfState.Serial, &tfState.Raw, &tfState.Lock, &tfState.LockID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

// Store stores the terraform state specified by ident
func (tfState *TfState) Store(data []byte) error {
	if err := json.Unmarshal(data, &tfState); err != nil {
		return err
	}
	tfState.Raw = string(data)

	q := "INSERT OR REPLACE INTO `tf_state` (`ident`, `version`, `serial`, `data`) VALUES ($1, $2, $3, $4);"

	if _, err := sqlite3.Exec(q, &tfState.Ident, &tfState.Version, &tfState.Serial, &tfState.Raw); err != nil {
		return err
	}

	return nil
}

// Lock locks a certain state according to the given flag
func Lock(ident string, lock bool, lockID string) error {
	query := "UPDATE tf_state SET `lock` = $1, `lock_id` = $2 WHERE `ident` = $3"
	if lock {
		if _, err := sqlite3.Exec(query, "1", lockID, ident); err != nil {
			return err
		}
	} else {
		if _, err := sqlite3.Exec(query, "0", lockID, ident); err != nil {
			return err
		}
	}
	return nil
}
