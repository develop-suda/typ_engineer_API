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

	// HTTPリクエストパラメータは文字型で取得されるため、数値型に変換する
	intQuantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		logs.WriteLog(err.Error(), def.ERROR)
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
	if result = selectItems.GetTypWords(db, values); &result == nil {
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

	db := connect.DbConnect()
	defer db.Close()
	if result = selectItems.GetTypes(db); result == nil {
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

	db := connect.DbConnect()
	defer db.Close()
	if result = selectItems.GetPartsOfSpeeches(db); result == nil {
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

	// TODO トランザクションはする
	// TODO connect.DbConnect()の帰り値がnilの場合の処理を追加する 全部の処理に追加する
	db := connect.DbConnect()
	defer db.Close()
	if db == nil {
		fmt.Println(err)
		return
	}

	if err = insert.CreateUser(db, values); err != nil { return }
	if userId = selectItems.MatchUserPassword(db, userMatcInfo); userId == "" { return }
	if err = insert.InsertTypWordInformation(db, userId); err != nil { return }
	if err = insert.InsertTypAlphabetInformation(db, userId); err != nil { return }
	if err = login.InsertLoginData(db, userId); err != nil { return }
	if loginData = login.CreateToken(userId); &loginData == nil { return }

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

	values := def.UserMatchInfo{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	db := connect.DbConnect()
	defer db.Close()

	tx, err := db.Begin()
	// deferが二つあった場合下の方から実行される
	defer db.Close()
	defer tx.Commit()
	if err != nil {
		panic(err)
	}

	// ユーザIDとパスワードが一致するか確認
	// 一致した場合はユーザIDを取得
	if userId = selectItems.MatchUserPassword(db, values); userId == "" {
		return
	}

	// ユーザIDを取得できた場合ログインデータをDBにインサート
	if err = login.InsertLoginData(db, userId); err != nil{
		return
	}

	// ユーザIDを取得できた場合jwtトークンを発行
	if loginData = login.CreateToken(userId); &loginData == nil {
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

	userId := def.UserIdStruct{User_id: r.FormValue("userId")}

	db := connect.DbConnect()
	defer db.Close()
	if err := logout.UpdateLogoutData(db, userId); err != nil {
		return
	}

}

func UpdateTypeInfoHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	var typWordInfos []def.TypWordInfo
	var typAlphaInfos []def.TypAlphabetInfo	

	values := map[string]string{
		 "typWordInfo":        r.FormValue("typWordInfo"),
		 "typAlphabetInfo":    r.FormValue("typAlphaInfo"),
	}

	userId := def.UserIdStruct{User_id: r.FormValue("userId")}

	// jsonを構造体に変換できるかどうか確認
 	err := json.Unmarshal([]byte(values["typWordInfo"]), &typWordInfos)
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
	
	db := connect.DbConnect()
	defer db.Close()
	
	//DBから単語情報を取得
	if wordDetails = selectItems.GetWordDetail(db); wordDetails== nil {
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

// TODO 関数名もろもろ変
// 関数名が下だけど
// 別の構造体にTypWordInfoという名前が使われてるためTypCountに近しい関数名にしよう
func GetWordTypInfoHandler(w http.ResponseWriter, r *http.Request){

	var typWordInfos []def.TypCount

	userId := def.UserIdStruct{User_id: r.FormValue("userId")}

	db := connect.DbConnect()
	defer db.Close()

	//DBから単語入力情報を取得
	if typWordInfos = selectItems.GetWordTypInfo(db, userId); typWordInfos == nil {
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

	db := connect.DbConnect()
	defer db.Close()

	//DBからユーザ情報を取得
	// 単語の入力成功、失敗回数を取得
	if myPageData.WordTypInfoSum = selectItems.GetWordTypInfoSum(db, userId); &myPageData.WordTypInfoSum == nil { return }

	// アルファベットの入力成功、失敗回数を取得
	if myPageData.AlphabetTypInfoSum = selectItems.GetAlphabetTypInfoSum(db, userId); &myPageData.AlphabetTypInfoSum == nil { return }

	// 単語の入力成功、失敗回数のランキングを取得
	if myPageData.WordCountRanking = selectItems.GetWordCountRanking(db, userId); &myPageData.WordCountRanking == nil { return }
	if myPageData.WordMissCountRanking = selectItems.GetWordMissCountRanking(db, userId); &myPageData.WordMissCountRanking == nil { return }

	// アルファベットの入力成功、失敗回数のランキングを取得
	if myPageData.AlphabetCountRanking = selectItems.GetAlphabetCountRanking(db, userId); &myPageData.AlphabetCountRanking == nil { return }
	if myPageData.AlphabetMissCountRanking = selectItems.GetAlphabetMissCountRanking(db, userId); &myPageData.AlphabetMissCountRanking == nil { return }

	//DBの取得結果をjsonに変換
	json, err := json.Marshal(myPageData)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	w.Write(json)

}