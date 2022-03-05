package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "bytes"
	"encoding/json"
	"log"
	"net/http"
    "net/url"
    "pkg/db/connect"
)

type word struct {
    Word    string  `json:"word"`
    Parts_of_speech string `json:"parts_of_speech"`
    Discription string `json:"Discription"`
}

type level struct {
    Level string `json:"level"`
}

func main() {
    handler1 := func(w http.ResponseWriter, r *http.Request) {

        urlParams := r.RequestURI

        db := dbConnect()
        result := dbSelect(db,urlParams)
        defer db.Close()

        var buf bytes.Buffer
        enc := json.NewEncoder(&buf)
        if err := enc.Encode(&result.Value); err != nil {
            log.Fatal(err)
        }

        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

        _, err := fmt.Fprint(w, buf.String())
        if err != nil {
            return
        }
	}

    handler2 := func(w http.ResponseWriter, r *http.Request) {

        urlParams := r.RequestURI

        db := dbConnect()
        result := getLevels(db,urlParams)
        defer db.Close()

        var buf bytes.Buffer
        enc := json.NewEncoder(&buf)
        if err := enc.Encode(&result.Value); err != nil {
            log.Fatal(err)
        }

        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

        _, err := fmt.Fprint(w, buf.String())
        if err != nil {
            return
        }
	}

	http.HandleFunc("/", handler1)
    http.HandleFunc("/levels", handler2)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func dbSelect(db *gorm.DB, urlParams string) *gorm.DB {

    // 複数件取得する場合、構造体を配列にする
    var words []word

    sql := "SELECT word, parts_of_speech, discription FROM words WHERE 1 = 1"
    u, err := url.Parse(urlParams)
    if err != nil {
        log.Fatal(err)
    }

    for key, values := range u.Query() {
        if key != "limit" {
            sql += " AND " + key + " = '" + values[0] + "'"
        } else {
            sql += " LIMIT " + values[0]
        }
    }

    result := db.Raw(sql).Find(&words)
    return result
}

func getLevels(db *gorm.DB, urlParams string) *gorm.DB {

    // 複数件取得する場合、構造体を配列にする
    var levels []level

    sql := "SELECT level FROM words GROUP BY level"

    result := db.Raw(sql).Find(&levels)
    return result
}