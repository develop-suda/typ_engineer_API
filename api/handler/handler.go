package handler

import (
	// "encoding/json"
	"net/http"
	"database/sql"
	"fmt"

	// insert "github.com/develop-suda/typ_engineer_API/api/insert"
	// selectItems "github.com/develop-suda/typ_engineer_API/api/select"
	// update "github.com/develop-suda/typ_engineer_API/api/update"
	connect "github.com/develop-suda/typ_engineer_API/internal/db"
	// login "github.com/develop-suda/typ_engineer_API/api/login"
	// logout "github.com/develop-suda/typ_engineer_API/api/logout"
	"log"
	logs "github.com/develop-suda/typ_engineer_API/internal/log"
	"github.com/go-sql-driver/mysql"


	def "github.com/develop-suda/typ_engineer_API/common"
)

func TypWordSelectHandler(w http.ResponseWriter, r *http.Request) {

	// values := map[string]string{
	// 	"type":            r.FormValue("type"),
	// 	"parts_of_speech": r.FormValue("parts_of_speech"),
	// 	"alphabet":        r.FormValue("alphabet"),
	// 	"quantity":        r.FormValue("quantity"),
	// }

	// db := connect.DbConnect()
	// result := selectItems.GetTypWords(db, values)
	// defer db.Close()

	// json, err := json.Marshal(result)
	// if err != nil {
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	// w.Write(json)
}

func GetTypeHandler(w http.ResponseWriter, r *http.Request) {

	// db := connect.DbConnect()
	// result := selectItems.GetTypes(db)
	// defer db.Close()

	// //DBの取得結果をjsonに変換
	// json, err := json.Marshal(result)
	// if err != nil {
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	// w.Write(json)
}

func GetPartsOfSpeechHandler(w http.ResponseWriter, r *http.Request) {

	// db := connect.DbConnect()
	// result := selectItems.GetPartsOfSpeeches(db)
	// defer db.Close()

	// //DBの取得結果をjsonに変換
	// json, err := json.Marshal(result)
	// if err != nil {
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	// w.Write(json)
}

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {

	values := map[string]string{
		"last_name":  r.FormValue("last_name"),
		"first_name": r.FormValue("first_name"),
		"email":      r.FormValue("email"),
		"password":   r.FormValue("password"),
	}


	var err error

	s := def.Service{}

	s.Tx.DB, err = connect.DbConnect(s.Tx.DB)
	if err != nil {
		fmt.Println(err)
	}
	defer s.Tx.Close()

	// ユーザ登録時の処理をトランザクションで実行
	s.UserRegistTran(values)

	// insert.CreateUser(db, values)
	// userId := selectItems.MatchUserPassword(db, values)
	// insert.InsertTypWordInformation(db, userId)
	// insert.InsertTypAlphabetInformation(db, userId)
	// login.InsertLoginData(db, userId)

	// loginData := login.CreateToken(userId)

	//loginDataをjsonに変換
	// json, err := json.Marshal(loginData)
	if err != nil {
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	// w.Write(json)

}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	// w.Header().Set("Access-Control-Allow-Headers", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	// values := map[string]string{
	// 	"email":    r.FormValue("email"),
	// 	"password": r.FormValue("password"),
	// }

	// db := connect.DbConnect()

	// tx, err := db.Begin()
	// if err != nil {
	// 	panic(err)
	// }

	// userId := selectItems.MatchUserPassword(db, values)
	// login.InsertLoginData(db, userId)
	// tx.Commit()

	// defer db.Close()

	// // ユーザIDを取得できた場合jwtトークンを発行
	// if userId != "" {

	//     loginData := login.CreateToken(userId)

	// 	//loginDataをjsonに変換
	// 	json, err := json.Marshal(loginData)
	// 	if err != nil {
	// 		return
	// 	}

	// 	w.Write(json)

	// }

}

func UserLogoutHandler(w http.ResponseWriter, r *http.Request) {
		
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	// w.Header().Set("Access-Control-Allow-Headers", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	// userId := r.FormValue("userId")

	// db := connect.DbConnect()
	// logout.UpdateLogoutData(db, userId)
	// defer db.Close()

}

func PostTypeWordInfoHandler(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	// w.Header().Set("Access-Control-Allow-Headers", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	// values := map[string]string{
	// 	 "userId":             r.FormValue("userId"),
	// 	 "typWordInfo":        r.FormValue("typWordInfo"),
	// }

	// var typWordInfo []def.TypWordInfo

	// // sql := def.UPDATE_TYP_WORD_INFO_SQL

	// // userId := values["userId"]
	// temp := values["typWordInfo"]
 	// json.Unmarshal([]byte(temp), &typWordInfo)

	// db := connect.DbConnect()
	// update.UpdateTypWordInfo(db, values)
	// // update.UpdateTypAlphabetInfo(db, values)
	// defer db.Close()

}

// txAdminはトランザクション制御するための構造体
// Transaction はトランザクションを制御するメソッド
//  アプリケーション開発者が本メソッドを使って、DMLのクエリーを発行する
func (t *def.TxAdmin) Transaction(f func() (err error)) error {
	// db,err := connect.StructDbConnect(t.DB)
	tx, err := t.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err = f(); err != nil {
		return fmt.Errorf("transaction query failed: %w", err)
	}
	return tx.Commit()
}

func (s *def.Service) UserRegistTran(values map[string]string) error {
	
	// user登録する関数
	userRegist := func() error {
		_, err := s.Tx.Exec(def.INSERT_USER_SQL, values["first_name"], values["last_name"], values["email"], values["password"])
		if err != nil {
			if mysqlErr, ok := err.(*mysql.MySQLError); ok {
				logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+def.INSERT_USER_SQL, def.ERROR)
			}
			log.Fatal(err)
		}
		return nil
	}

	all := func() error {
		userRegist()
		// MatchUserPassword()
		s.Test(values)
		return nil
	}

	return s.Tx.Transaction(all)
}

func (s *Service)Test(values map[string]string ) string {
		
	var user_id string
	var err error

	sql := "SELECT LPAD(user_id,8,0) FROM users WHERE email = ? AND password = ?"

	result := s.Tx.QueryRow(sql, values["email"], values["password"])
	if err = result.Err(); err != nil {
		fmt.Println(err)
	}
	
	err = result.Scan(&user_id)
	if err != nil {
		log.Fatal(err)
	}

	return user_id
}
