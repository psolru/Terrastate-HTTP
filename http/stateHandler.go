package http

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/psolru/terrastate-http/statemanager"
	"io/ioutil"
	"net/http"
)

const LOCKED = 1

func stateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tfState := &statemanager.TfState{Ident: vars["ident"]}
	err := tfState.Get()
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if r.Method == "GET" && err == sql.ErrNoRows {
		http.NotFound(w, r)
	} else if r.Method == "POST" {
		if tfState.Lock == LOCKED && tfState.LockID != r.URL.Query().Get("ID") {
			http.Error(w, "", http.StatusLocked)
			return
		}
		var bytes []byte
		if bytes, err = ioutil.ReadAll(r.Body); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		if err = tfState.Store(bytes); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}

	_, _ = w.Write([]byte(tfState.Raw))
}
