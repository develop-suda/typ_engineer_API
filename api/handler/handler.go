package handler

import (
	"encoding/json"
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

	// パラメータをvaluesに格納
	values := def.TypWordSelect{
		Word_type:       r.FormValue("type"),
		Parts_of_speech: r.FormValue("parts_of_speech"),
		Alphabet:        r.FormValue("alphabet"),
		Quantity:        intQuantity,
	}

	// DBに接続
	db := connect.DbConnect()
	if db == nil {
		return
	}

	// DB接続を閉じる
	defer db.Close()

	// DBからデータを取得
	if result, err = selectItems.GetTypWords(db, values); err != nil {
		return
	}

	// DBの取得結果をjsonに変換
	json, err := json.Marshal(result)
	if err != nil {
		return
	}

	// ヘッダーを設定
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	// jsonを返す
	w.Write(json)
}

func GetTypeHandler(w http.ResponseWriter, r *http.Request) {

	var result []def.WordType
	var err error

	// DBに接続
	db := connect.DbConnect()
	if db == nil {
		return
	}

	// DB接続を閉じる
	defer db.Close()

	// DBからデータを取得
	if result, err = selectItems.GetTypes(db); err != nil {
		return
	}

	//DBの取得結果をjsonに変換
	json, err := json.Marshal(result)
	if err != nil {
		return
	}

	// ヘッダーを設定
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	// jsonを返す
	w.Write(json)
}

func GetPartsOfSpeechHandler(w http.ResponseWriter, r *http.Request) {

	var result []def.PartsOfSpeech
	var err error

	// DBに接続
	db := connect.DbConnect()
	if db == nil {
		return
	}
	// DB接続を閉じる
	defer db.Close()

	// DBからデータを取得
	if result, err = selectItems.GetPartsOfSpeeches(db); err != nil {
		return
	}

	//DBの取得結果をjsonに変換
	json, err := json.Marshal(result)
	if err != nil {
		return
	}

	// ヘッダーを設定
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	// jsonを返す
	w.Write(json)
}

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {

	var err error
	var userId string
	var loginData def.LoginData

	// パラメータをvaluesに格納
	values := def.UserRegisterInfo{
		Last_name:  r.FormValue("last_name"),
		First_name: r.FormValue("first_name"),
		Email:      r.FormValue("email"),
		Password:   r.FormValue("password"),
	}

	// パスワードが一致か確認するために取得
	userMatchInfo := def.UserMatchInfo{
		Email:    values.Email,
		Password: values.Password,
	}
	
	// DBに接続
	db := connect.DbConnect()
	if db == nil {
		return
	}

	// トランザクション開始
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

	// ユーザ登録
	if err = insert.CreateUser(tx, values); err != nil { return }
	// パラメータにあるemailとpasswordが一致するユーザIDを取得
	if userId, err = selectItems.TranMatchUserPassword(tx, userMatchInfo); err != nil { return }
	// 単語の初期情報を登録
	if err = insert.InsertTypWordInformation(tx, userId); err != nil { return }
	// アルファベットの初期情報を登録
	if err = insert.InsertTypAlphabetInformation(tx, userId); err != nil { return }
	// ログイン情報を登録
	if err = login.TranInsertLoginData(tx, userId); err != nil { return }
	// ユーザIDを取得できた場合jwtトークンを発行
	if loginData, err = login.CreateToken(userId); err != nil { return }

	// loginDataをjsonに変換
	json, err := json.Marshal(loginData)
	if err != nil {
		return
	}

	// ヘッダーを設定
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	// jsonを返す
	w.Write(json)

}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {

	var userId string
	var loginData def.LoginData
	var err error

	// パラメータをvaluesに格納
	values := def.UserMatchInfo{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// DBに接続
	db := connect.DbConnect()
	if err != nil {
		return
	}
	
	// DB接続を閉じる
	defer db.Close()

	// パラメータにあるemailとpasswordが一致するユーザIDを取得
	if userId, err = selectItems.MatchUserPassword(db, values); err != nil { return }
	// ログイン情報を登録
	if err = login.InsertLoginData(db, userId); err != nil { return }
	// ユーザIDを取得できた場合jwtトークンを発行
	if loginData, err = login.CreateToken(userId); err != nil { return }

	//loginDataをjsonに変換
	json, err := json.Marshal(loginData)
	if err != nil {
		return
	}

	// ヘッダーを設定
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

	// jsonを返す
	w.Write(json)

}

func UserLogoutHandler(w http.ResponseWriter, r *http.Request) {

	var err error

	// パラメータをuserIdに格納
	userId := def.UserIdStruct{User_id: r.FormValue("userId")}

	// DBに接続
	db := connect.DbConnect()
	if err != nil {
		return
	}
	// DB接続を閉じる
	defer db.Close()

	// ログアウト情報を登録
	if err = logout.UpdateLogoutData(db, userId); err != nil { return }

	// ヘッダーを設定
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

}

func UpdateTypeInfoHandler(w http.ResponseWriter, r *http.Request) {

	var typWordInfos []def.TypWordInfo
	var typAlphaInfos []def.TypAlphabetInfo	
	var err error

	// パラメータをvaluesに格納
	values := map[string]string{
		 "typWordInfo":        r.FormValue("typWordInfo"),
		 "typAlphabetInfo":    r.FormValue("typAlphaInfo"),
	}

	// パラメータをuserIdに格納
	userId := def.UserIdStruct{User_id: r.FormValue("userId")}

	// 単語の入力情報jsonを構造体に変換できるかどうか確認
 	err = json.Unmarshal([]byte(values["typWordInfo"]), &typWordInfos)
	if err != nil {
		return
	}

	// アルファベット入力情報jsonを構造体に変換できるかどうか確認
	err = json.Unmarshal([]byte(values["typAlphabetInfo"]), &typAlphaInfos)
	if err != nil {
		return
	}

	// DBに接続
	db := connect.DbConnect()
	if db == nil {
		return
	}

	// トランザクション開始
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
	
	// 単語の入力情報を更新
	if err = update.UpdateTypWordInfo(tx, typWordInfos, userId); err != nil { return }
	// アルファベットの入力情報を更新
	if err = update.UpdateTypAlphabetInfo(tx, typAlphaInfos, userId); err != nil { return }

	// ヘッダーを設定
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

}

func GetWordDetailHandler(w http.ResponseWriter, r *http.Request) {

	var wordDetails []def.WordDetail
	var err error
	
	// DBに接続
	db := connect.DbConnect()
	if err != nil {
		return
	}

	// DB接続を閉じる
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

	// ヘッダーを設定
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	// jsonを返す
	w.Write(json)

}

func GetTypCountAllWordHandler(w http.ResponseWriter, r *http.Request){

	var typWordInfos []def.TypCount
	var err error

	// パラメータをuserIdに格納
	userId := def.UserIdStruct{User_id: r.FormValue("userId")}

	// DBに接続
	db := connect.DbConnect()
	if err != nil {
		return
	}
	
	// DB接続を閉じる
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

	// ヘッダーを設定
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	// jsonを返す
	w.Write(json)
}

func GetMyPageDataHandler(w http.ResponseWriter, r *http.Request) {

	var myPageData def.MyPageData
	var err error

	// パラメータをuserIdに格納
	userId := def.UserIdStruct{User_id: r.FormValue("userId")}

	// DBに接続
	db := connect.DbConnect()
	if err != nil {
		return
	}

	// DB接続を閉じる
	defer db.Close()

	// DBからユーザ情報を取得
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

	// DBの取得結果をjsonに変換
	json, err := json.Marshal(myPageData)
	if err != nil {
		return
	}

	// ヘッダーを設定
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	// jsonを返す
	w.Write(json)

}