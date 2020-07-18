package http

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/psolru/terrastate-http/statemanager"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func lockHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	lockID, err := extractLockIDByBody(r.Body)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

	if err := statemanager.Lock(vars["ident"], true, lockID); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "", http.StatusNotFound)
		}
		http.Error(w, "", http.StatusInternalServerError)
	}
	_, _ = w.Write([]byte(""))
}

func unlockHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := statemanager.Lock(vars["ident"], false, ""); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "", http.StatusNotFound)
		}
		http.Error(w, "", http.StatusInternalServerError)
	}
	_, _ = w.Write([]byte(""))
}

func extractLockIDByBody(body io.ReadCloser) (string, error) {
	bytes, _ := ioutil.ReadAll(body)

	var objmap map[string]*json.RawMessage
	if err := json.Unmarshal(bytes, &objmap); err != nil {
		return "", err
	}

	byteID, _ := objmap["ID"].MarshalJSON()
	return strings.Trim(string(byteID), "\""), nil
}
