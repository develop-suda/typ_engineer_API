package def

import (
	"strconv"
	"fmt"
	"database/sql"

	"github.com/go-sql-driver/mysql"

	logs "github.com/develop-suda/typ_engineer_API/internal/log"
	def "github.com/develop-suda/typ_engineer_API/common"
)

type tempStructWord struct {
	Word string
}

// ユーザ情報を登録する関数
func CreateUser(tx *sql.Tx, values def.UserRegisterInfo) error {
	logs.WriteLog("CreateUser開始", nil, def.NORMAL)

	var err error

    sql := def.GetInsertUserSQL()

	err = values.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), values, def.ERROR)
		return err
	}

	//SQL実行
	_, err = tx.Query(sql, values.First_name, values.Last_name, values.Email, values.Password)
	
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, values, def.ERROR)
		}
		return err
	}

	logs.WriteLog("CreateUser正常終了", nil, def.NORMAL)
	return nil
}

// wordのタイピング情報の更新先を登録する関数
func InsertTypWordInformation(tx *sql.Tx, userId string) error {
	logs.WriteLog("InsertTypeWordInfo開始", nil, def.NORMAL)

	var words []tempStructWord

	sql := def.GetWordUniqueSQL()

	// 重複しない全単語を取得
	selectWords, err := tx.Query(sql)
	if err != nil {
		// TODO 調べる
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, userId, def.ERROR)
		}
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return err
	}

	// 取得した単語を配列に格納
	for selectWords.Next() {
		word := tempStructWord{}
		// TODO 調べる
		if err := selectWords.Scan(&word.Word); err != nil {
			logs.WriteLog(err.Error(), 
				tempStructWord{Word: word.Word},def.ERROR)
			return err
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
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, userId, def.ERROR)
		}
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return err
	}

	// 登録先のテーブル、登録件数をログに出力
	if rows,err := result.RowsAffected(); err == nil {
		logs.WriteLog("テーブル名 : typing_word_informations ,インサート件数 : " + strconv.FormatInt(rows,10), sql, def.NORMAL)
	}


	logs.WriteLog("InsertTypWordInfo終了", nil,def.NORMAL)
	return nil
}

// アルファベットのタイピング情報の更新先を登録する関数
func InsertTypAlphabetInformation(tx *sql.Tx, userId string) error {
	logs.WriteLog("InsertTypAlphabetInfo開始", nil, def.NORMAL)

	sql := def.INSERT_TYPING_ALPHABET_INFORMATIONS_SQL

	//　アルファベットのタイピング情報の更新先を登録する
	// unicodeはa~zまでのアルファベットのコード
	for unicode := 97; unicode <= 122; unicode++ {
		sql += "(" + userId + ", '" + string(unicode) + "',0,0,cast(now() as datetime),cast(now() as datetime),0),"
	}

	sql = sql[:len(sql)-1]

	// アルファベットのタイピング情報の更新先を登録する
	result, err := tx.Exec(sql)
	if err != nil {
		// TODO 調べる
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, userId, def.ERROR)
		}
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return err
	}

	// 登録先のテーブル、登録件数をログに出力
	if rows,err := result.RowsAffected(); err == nil {
		logs.WriteLog("テーブル名 : typing_alphabet_informations ,インサート件数 : " + strconv.FormatInt(rows,10), sql, def.NORMAL)
	}

	logs.WriteLog("InsertTypAlphabetInfo終了", nil, def.NORMAL)
	return nil
}