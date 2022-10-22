package insert

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/go-sql-driver/mysql"

	def "github.com/develop-suda/typ_engineer_API/common"
	logs "github.com/develop-suda/typ_engineer_API/internal/log"
)

type tempStructWord struct {
	Word string
}

// ユーザ情報を登録する関数
func CreateUser(values map[string]string) {
	// logs.WriteLog("CreateUser開始", def.NORMAL)
    // sql := def.INSERT_USER_SQL

	// //SQL実行
	// _, err := s.Tx.DB.Exec(sql, values["first_name"], values["last_name"], values["email"], values["password"])
	
	// if err != nil {
	// 	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
	// 		logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
	// 	}
	// 	log.Fatal(err)
	// }

	// logs.WriteLog("CreateUser正常終了", def.NORMAL)
	// return
}

// wordのタイピング情報の更新先を登録する関数
func InsertTypWordInformation(tx *sql.DB, userId string) {
	logs.WriteLog("InsertTypeWordInfo開始",def.NORMAL)

	var words []tempStructWord

	sql := def.GET_WORD_UNIQUE_SQL

	// 重複しない全単語を取得
	selectWords, err := tx.Query(sql)
	if err != nil {
		// TODO 調べる
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	// 取得した単語を配列に格納
	for selectWords.Next() {
		word := tempStructWord{}
		// TODO 調べる
		if err := selectWords.Scan(&word.Word); err != nil {
			log.Fatal(err)
		}
		words = append(words, word)
	}

	sql = def.INSERT_TYPING_WORD_INFORMATIONS_SQL

	// 取得した単語を元にINSERT文を作成
	for _,value := range words {
		sql += "(" + userId + ", '" + value.Word + "',0,0,cast(now() as datetime),cast(now() as datetime),0),"
	}

	sql = sql[:len(sql)-1]

	// wordのタイピング情報の更新先を登録する
	result, err := tx.Exec(sql)
	if err != nil {
		// TODO 調べる
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	// 登録先のテーブル、登録件数をログに出力
	if rows,err := result.RowsAffected(); err == nil {
		logs.WriteLog("テーブル名 : typing_word_informations ,インサート件数 : " + strconv.FormatInt(rows,10), def.NORMAL)
	}


	logs.WriteLog("InsertTypWordInfo終了", def.NORMAL)
}

// アルファベットのタイピング情報の更新先を登録する関数
func InsertTypAlphabetInformation(db *sql.DB, userId string) {
	logs.WriteLog("InsertTypAlphabetInfo開始",def.NORMAL)

	sql := def.INSERT_TYPING_ALPHABET_INFORMATIONS_SQL

	//　アルファベットのタイピング情報の更新先を登録する
	// unicodeはa~zまでのアルファベットのコード
	for unicode := 97; unicode <= 122; unicode++ {
		sql += "(" + userId + ", '" + string(unicode) + "',0,0,cast(now() as datetime),cast(now() as datetime),0),"
	}

	sql = sql[:len(sql)-1]

	// アルファベットのタイピング情報の更新先を登録する
	result, err := db.Exec(sql)
	if err != nil {
		// TODO 調べる
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	// 登録先のテーブル、登録件数をログに出力
	if rows,err := result.RowsAffected(); err == nil {
		logs.WriteLog("テーブル名 : typing_alphabet_informations ,インサート件数 : " + strconv.FormatInt(rows,10), def.NORMAL)
	}

	logs.WriteLog("InsertTypAlphabetInfo終了", def.NORMAL)
}