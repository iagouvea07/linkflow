package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

var result string

func HandlerRedirect(db *sql.DB, rdb *redis.Client, encode string, w http.ResponseWriter, r *http.Request) {
	result, err := rdb.Get(ctx, "encode:" + encode[1:]).Result()

	if err != nil {
		err := db.QueryRow("SELECT url FROM addresses WHERE encoded = ?", encode[1:]).Scan(&result)

		if err != nil {
			fmt.Println(err)
			return
		}

		err = rdb.Set(ctx, "encode:" + encode[1:], result, 1 * time.Hour).Err()

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if !strings.HasPrefix(result, "http://") || !strings.HasPrefix(result, "https://") {
		result = "https://" + result
	}

	http.Redirect(w, r, result, http.StatusSeeOther)
}