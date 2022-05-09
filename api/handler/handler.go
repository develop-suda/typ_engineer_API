package handler

import (
	"encoding/json"
	"net/http"

	// userRegist "github.com/develop-suda/typ_engineer_API/api/insert"
	selectItems "github.com/develop-suda/typ_engineer_API/api/select"
	connect "github.com/develop-suda/typ_engineer_API/internal/db"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TypWordSelectHandler(w http.ResponseWriter, r *http.Request) {

	// ↓でGetPostのifを作る
	// test := http.MethodPost
	// fmt.Print(test)

	values := map[string]string{
		"type":            r.FormValue("type"),
		"parts_of_speech": r.FormValue("parts_of_speech"),
		"alphabet":        r.FormValue("alphabet"),
		"quantity":        r.FormValue("quantity"),
	}

	db := connect.DbConnect()
	result := selectItems.GetTypWords(db, values)
	defer db.Close()

	json, err := json.Marshal(result)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	w.Write(json)
}

func GetTypeHandler(w http.ResponseWriter, r *http.Request) {

	db := connect.DbConnect()
	result := selectItems.GetTypes(db)
	defer db.Close()

	//DBの取得結果をjsonに変換
	json, err := json.Marshal(result)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	w.Write(json)
}

func GetPartsOfSpeechHandler(w http.ResponseWriter, r *http.Request) {

	db := connect.DbConnect()
	result := selectItems.GetPartsOfSpeeches(db)
	defer db.Close()

	//DBの取得結果をjsonに変換
	json, err := json.Marshal(result)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	w.Write(json)
}

// func UserRegistHandler(w http.ResponseWriter, r *http.Request) {

// 	values := map[string]string{
// 		"user_name":    r.FormValue("user_name"),
// 		"email":        r.FormValue("email"),
// 		"phone_number": r.FormValue("phone_number"),
// 	}

// 	db := connect.DbConnect()
// 	userRegist.InsertUser(db, values)
// 	result := selectItems.GetUsers(db)
// 	defer db.Close()

// 	var buf bytes.Buffer
// 	json.Conversion(result, &buf)

// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

// 	_, err := fmt.Fprint(w, buf.String())
// 	if err != nil {
// 		return
// 	}
// }
