package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

var result string

func HandlerRedirect(db *sql.DB, encode string, w http.ResponseWriter, r *http.Request) {
	err := db.QueryRow("SELECT url FROM addresses WHERE encoded = ?", encode[1:]).Scan(&result)

	if err != nil {
		fmt.Println(err)
		return
	}

	if !strings.HasPrefix(result, "http://") || !strings.HasPrefix(result, "https://") {
		result = "https://" + result
	}

	fmt.Println(result)
	http.Redirect(w, r, result, http.StatusSeeOther)
}