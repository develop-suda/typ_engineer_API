package handler

import (
	"net/http"
	"fmt"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "bytes"
	"github.com/develop-suda/typ_engineer_API/internal/db"
    "github.com/develop-suda/typ_engineer_API/internal/conversion"
    "github.com/develop-suda/typ_engineer_API/api/select"
)


func SelectWordsHandler(w http.ResponseWriter, r *http.Request) {

    urlParams := r.RequestURI

    db := connect.DbConnect()
    result := selectItems.DbSelect(db,urlParams)
    defer db.Close()

    var buf bytes.Buffer
    json.Conversion(result, &buf)


    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

    _, err := fmt.Fprint(w, buf.String())
    if err != nil {
        return
    }
}

func GetLevelsHandler(w http.ResponseWriter, r *http.Request) {

    urlParams := r.RequestURI

    db := connect.DbConnect()
    result := selectItems.GetLevels(db,urlParams)
    defer db.Close()

    var buf bytes.Buffer
    json.Conversion(result, &buf)

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

    _, err := fmt.Fprint(w, buf.String())
    if err != nil {
        return
    }
}