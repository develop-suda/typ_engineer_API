package handler

import (
	"encoding/json"
	"net/http"
	"fmt"

	insert "github.com/develop-suda/typ_engineer_API/api/insert"
	selectItems "github.com/develop-suda/typ_engineer_API/api/select"
	update "github.com/develop-suda/typ_engineer_API/api/update"
	connect "github.com/develop-suda/typ_engineer_API/internal/db"
	login "github.com/develop-suda/typ_engineer_API/api/login"
	logout "github.com/develop-suda/typ_engineer_API/api/logout"
	def "github.com/develop-suda/typ_engineer_API/common"
)

func TypWordSelectHandler(w http.ResponseWriter, r *http.Request) {

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

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {

	values := map[string]string{
		"last_name":  r.FormValue("last_name"),
		"first_name": r.FormValue("first_name"),
		"email":      r.FormValue("email"),
		"password":   r.FormValue("password"),
	}


	var err error

	//トランザクションはする
	db := connect.DbConnect()
	if err != nil {
		fmt.Println(err)
	}

	insert.CreateUser(db, values)
	userId := selectItems.MatchUserPassword(db, values)
	insert.InsertTypWordInformation(db, userId)
	insert.InsertTypAlphabetInformation(db, userId)
	login.InsertLoginData(db, userId)

	loginData := login.CreateToken(userId)

	// loginDataをjsonに変換
	json, err := json.Marshal(loginData)
	if err != nil {
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	w.Write(json)

}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	values := map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	}

	db := connect.DbConnect()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	userId := selectItems.MatchUserPassword(db, values)
	login.InsertLoginData(db, userId)
	tx.Commit()

	defer db.Close()

	// ユーザIDを取得できた場合jwtトークンを発行
	if userId != "" {

	    loginData := login.CreateToken(userId)

		//loginDataをjsonに変換
		json, err := json.Marshal(loginData)
		if err != nil {
			return
		}

		w.Write(json)

	}

}

func UserLogoutHandler(w http.ResponseWriter, r *http.Request) {
		
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	userId := r.FormValue("userId")

	db := connect.DbConnect()
	logout.UpdateLogoutData(db, userId)
	defer db.Close()

}

func UpdateTypeInfoHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	values := map[string]string{
		 "userId":             r.FormValue("userId"),
		 "typWordInfo":        r.FormValue("typWordInfo"),
		 "typAlphabetInfo":    r.FormValue("typAlphaInfo"),
	}

	var typWordInfo []def.TypWordInfo
	var typAlphaInfo []def.TypAlphabetInfo	

	// jsonを構造体に変換
 	err := json.Unmarshal([]byte(values["typWordInfo"]), &typWordInfo)
	if err != nil {
		fmt.Println(err)
	}

	// jsonを構造体に変換
	err = json.Unmarshal([]byte(values["typAlphabetInfo"]), &typAlphaInfo)
	if err != nil {
		fmt.Println(err)
	}

	// ToDo トランザクションにする
	// 入力した単語情報とアルファベット情報をDBに登録
	db := connect.DbConnect()
	update.UpdateTypWordInfo(db, values)
	update.UpdateTypAlphabetInfo(db, values)
	defer db.Close()

}
