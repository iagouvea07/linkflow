package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/speps/go-hashids/v2"
)

type Url struct {
	URL string `json:"url"`
}

type UrlEncoded struct {
	ID uint64 		`json:"id"`
	URL string  	`json:"url"`
	ENCODE string 	`json:"encode"`	
}

var current_position int
var mu sync.Mutex


func InitializePosition(db *sql.DB) error {
	err := db.QueryRow("SELECT current_position FROM position").Scan(&current_position)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Error: %w", err)
	}

	if current_position < 1 { 
		current_position = 1
	}

	return nil
}

func HandleEncode(db *sql.DB, url string, id uint64, w http.ResponseWriter) {
	encoded := genHashId(id)

    data := UrlEncoded{
        ID:     id,
        URL:    url,
        ENCODE: string(encoded), 
    }

	if _, err := db.Exec("INSERT INTO addresses (id, url, encoded) VALUES (?, ?, ?)", data.ID, data.URL, data.ENCODE); err != nil {
		fmt.Println(err)
		return
	}

    w.Header().Set("Content-Type", "application/json")

    err := json.NewEncoder(w).Encode(data)
    if err != nil {
        return
    }
}

func genHashId(id uint64) string {
	hd := hashids.NewData()
	hd.Salt = "fj4u9f6m5329c40u5"
	h, _ := hashids.NewWithData(hd)

	encoded, _ := h.Encode([]int{int(id)})
	return encoded
}


func HandleInsertUrl(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var data Url
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Erro ao processar JSON", http.StatusBadRequest)
        return
	}

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer tx.Rollback()

	err = tx.QueryRow("SELECT current_position FROM position FOR UPDATE").Scan(&current_position)
	if err != nil {
		http.Error(w, "Erro ao ler posição", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	current_position++

	if _, err = tx.Exec("UPDATE position SET current_position = ?", current_position); err != nil {
		fmt.Println(err)
		return
	}
	
	if err = tx.Commit(); err != nil {
		http.Error(w, "Erro ao confirmar transação", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	dataUrl := data.URL
	HandleEncode(db, dataUrl, uint64(current_position), w)
}