package http

import (
	"database/sql"
	"encoding/json"
	"github.com/psolru/terrastate-http/safeclose"
	"github.com/psolru/terrastate-http/sqlite3"
	"net/http"
)

type ListResponse struct {
	List map[int]string `json:"stateList"`
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := sqlite3.Query("SELECT ROWID as id, `ident` FROM tf_state;")
	defer safeclose.Close(rows)

	status := http.StatusOK
	var out []byte
	if err != nil {
		if err != sql.ErrNoRows {
			status = http.StatusInternalServerError
		}
	} else {
		list := make(map[int]string)
		for rows.Next() {
			var id int
			var ident string
			if err := rows.Scan(&id, &ident); err != nil {
				status = http.StatusInternalServerError
			} else {
				list[id] = ident
			}
		}
		err = rows.Err()
		if err != nil {
			status = http.StatusInternalServerError
		}

		out, err = json.Marshal(&list)
		if err != nil {
			status = http.StatusInternalServerError
		}
	}
	if status != http.StatusOK {
		http.Error(w, "{}", status)
		return
	}

	_, _ = w.Write(out)
}
