package handler

import (
	"bytes"
	"fmt"
	"net/http"

	userRegist "github.com/develop-suda/typ_engineer_API/api/insert"
	selectItems "github.com/develop-suda/typ_engineer_API/api/select"
	json "github.com/develop-suda/typ_engineer_API/internal/conversion"
	connect "github.com/develop-suda/typ_engineer_API/internal/db"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SelectWordsHandler(w http.ResponseWriter, r *http.Request) {

	urlParams := r.RequestURI

	//↓でGetPostのifを作る
	// test := http.MethodPost
	// fmt.Print(test)

	db := connect.DbConnect()
	result := selectItems.DbSelect(db, urlParams)
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
	result := selectItems.GetLevels(db, urlParams)
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

func UserRegistHandler(w http.ResponseWriter, r *http.Request) {

	values := map[string]string{
		"user_name":    r.FormValue("user_name"),
		"email":        r.FormValue("email"),
		"phone_number": r.FormValue("phone_number"),
	}

	db := connect.DbConnect()
	userRegist.InsertUser(db, values)
	result := selectItems.GetUsers(db)
	defer db.Close()

	var buf bytes.Buffer
	json.Conversion(result, &buf)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}
}
