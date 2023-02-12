package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	insert "github.com/develop-suda/typ_engineer_API/api/insert"
	login "github.com/develop-suda/typ_engineer_API/api/login"
	logout "github.com/develop-suda/typ_engineer_API/api/logout"
	selectItems "github.com/develop-suda/typ_engineer_API/api/select"
	update "github.com/develop-suda/typ_engineer_API/api/update"
	def "github.com/develop-suda/typ_engineer_API/common"
	connect "github.com/develop-suda/typ_engineer_API/internal/db"
	logs "github.com/develop-suda/typ_engineer_API/internal/log"
)

func TypWordSelectHandler(w http.ResponseWriter, r *http.Request) {

	var result []def.Word
	var err error

	// HTTPリクエストパラメータは文字型で取得されるため、数値型に変換する
	intQuantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		logs.WriteLog(err.Error(), r.FormValue("quantity"), def.ERROR)
		return
	}

	values := def.TypWordSelect{
		Word_type:       r.FormValue("type"),
		Parts_of_speech: r.FormValue("parts_of_speech"),
		Alphabet:        r.FormValue("alphabet"),
		Quantity:        intQuantity,
	}

	db := connect.DbConnect()
	defer db.Close()
	if result, err = selectItems.GetTypWords(db, values); err != nil {
		return
	}

	json, err := json.Marshal(result)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	w.Write(json)
}

func GetTypeHandler(w http.ResponseWriter, r *http.Request) {

	var result []def.WordType
	var err error

	db := connect.DbConnect()
	defer db.Close()
	if result, err = selectItems.GetTypes(db); err != nil {
		return
	}

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

	var result []def.PartsOfSpeech
	var err error

	db := connect.DbConnect()
	defer db.Close()
	if result, err = selectItems.GetPartsOfSpeeches(db); err != nil {
		return
	}

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

	var err error
	var userId string
	var loginData def.LoginData

	// HTTPリクエストパラメータを取得
	values := def.UserRegisterInfo{
		Last_name:  r.FormValue("last_name"),
		First_name: r.FormValue("first_name"),
		Email:      r.FormValue("email"),
		Password:   r.FormValue("password"),
	}

	// パスワードが一致か確認するために取得
	userMatcInfo := def.UserMatchInfo{
		Email:    values.Email,
		Password: values.Password,
	}
	
	db := connect.DbConnect()
	if err != nil {
		return
	}

	// トランザクション
	tx, err := db.Begin()
	if err != nil {
		return
	}

	// トランザクションのコミット、ロールバック
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err = insert.CreateUser(tx, values); err != nil { return }
	if userId, err = selectItems.TranMatchUserPassword(tx, userMatcInfo); err != nil { return }
	if err = insert.InsertTypWordInformation(tx, userId); err != nil { return }
	if err = insert.InsertTypAlphabetInformation(tx, userId); err != nil { return }
	if err = login.TranInsertLoginData(tx, userId); err != nil { return }
	if loginData, err = login.CreateToken(userId); err != nil { return }

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

	var userId string
	var loginData def.LoginData
	var err error

	values := def.UserMatchInfo{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	db := connect.DbConnect()
	if err != nil {
		panic(err)
	}
	
	defer db.Close()
	

	// ユーザIDとパスワードが一致するか確認
	// 一致した場合はユーザIDを取得
	if userId, err = selectItems.MatchUserPassword(db, values); err != nil {
		return
	}

	// ユーザIDを取得できた場合ログインデータをDBにインサート
	if err = login.InsertLoginData(db, userId); err != nil {
		return
	}

	// ユーザIDを取得できた場合jwtトークンを発行
	if loginData, err = login.CreateToken(userId); err != nil {
		return
	}

	//loginDataをjsonに変換
	json, err := json.Marshal(loginData)
	if err != nil {
		return
	}

	w.Write(json)

}

func UserLogoutHandler(w http.ResponseWriter, r *http.Request) {
		
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	var err error
	userId := def.UserIdStruct{User_id: r.FormValue("userId")}

	db := connect.DbConnect()
	defer db.Close()
	if err = logout.UpdateLogoutData(db, userId); err != nil {
		return
	}

}

func UpdateTypeInfoHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	var typWordInfos []def.TypWordInfo
	var typAlphaInfos []def.TypAlphabetInfo	
	var err error

	values := map[string]string{
		 "typWordInfo":        r.FormValue("typWordInfo"),
		 "typAlphabetInfo":    r.FormValue("typAlphaInfo"),
	}

	userId := def.UserIdStruct{User_id: r.FormValue("userId")}

	// jsonを構造体に変換できるかどうか確認
 	err = json.Unmarshal([]byte(values["typWordInfo"]), &typWordInfos)
	if err != nil {
		fmt.Println(err)
		return
	}

	// jsonを構造体に変換できるかどうか確認
	err = json.Unmarshal([]byte(values["typAlphabetInfo"]), &typAlphaInfos)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ToDo トランザクションにする
	// 入力した単語情報とアルファベット情報をDBに登録
	db := connect.DbConnect()
	defer db.Close()		
	if err = update.UpdateTypWordInfo(db, typWordInfos, userId); err != nil { return }
	if err = update.UpdateTypAlphabetInfo(db, typAlphaInfos, userId); err != nil { return }

}

func GetWordDetailHandler(w http.ResponseWriter, r *http.Request) {

	var wordDetails []def.WordDetail
	var err error
	
	db := connect.DbConnect()
	defer db.Close()
	
	//DBから単語情報を取得
	if wordDetails, err = selectItems.GetWordDetail(db); err != nil {
		return
	}

	//DBの取得結果をjsonに変換
	json, err := json.Marshal(wordDetails)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	w.Write(json)

}

func GetTypCountAllWordHandler(w http.ResponseWriter, r *http.Request){

	var typWordInfos []def.TypCount
	var err error

	userId := def.UserIdStruct{User_id: r.FormValue("userId")}

	db := connect.DbConnect()
	defer db.Close()

	//DBから単語入力情報を取得
	if typWordInfos, err = selectItems.GetWordTypInfo(db, userId); err != nil {
		return
	}

	//DBの取得結果をjsonに変換
	json, err := json.Marshal(typWordInfos)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	w.Write(json)
}

func GetMyPageDataHandler(w http.ResponseWriter, r *http.Request) {

	userId := def.UserIdStruct{User_id: r.FormValue("userId")}
	var myPageData def.MyPageData
	var err error

	db := connect.DbConnect()
	defer db.Close()

	//DBからユーザ情報を取得
	// 単語の入力成功、失敗回数を取得
	if myPageData.WordTypInfoSum, err = selectItems.GetWordTypInfoSum(db, userId); err != nil { return }

	// アルファベットの入力成功、失敗回数を取得
	if myPageData.AlphabetTypInfoSum, err = selectItems.GetAlphabetTypInfoSum(db, userId); err != nil { return }

	// 単語の入力成功、失敗回数のランキングを取得
	if myPageData.WordCountRanking, err = selectItems.GetWordCountRanking(db, userId); err != nil { return }
	if myPageData.WordMissCountRanking, err = selectItems.GetWordMissCountRanking(db, userId); err != nil { return }

	// アルファベットの入力成功、失敗回数のランキングを取得
	if myPageData.AlphabetCountRanking, err = selectItems.GetAlphabetCountRanking(db, userId); err != nil { return }
	if myPageData.AlphabetMissCountRanking, err = selectItems.GetAlphabetMissCountRanking(db, userId); err != nil { return }

	//DBの取得結果をjsonに変換
	json, err := json.Marshal(myPageData)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	w.Write(json)

}