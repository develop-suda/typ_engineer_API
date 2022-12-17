package update

import (
	"database/sql"
	"fmt"
	"log"
	"encoding/json"

	"github.com/go-sql-driver/mysql"

	logs "github.com/develop-suda/typ_engineer_API/internal/log"
	def "github.com/develop-suda/typ_engineer_API/common"
)

// wordのタイピング情報を更新する関数
func UpdateTypWordInfo(db *sql.DB, values map[string]string) {
	logs.WriteLog("UpdateTypWordInfo開始", def.NORMAL)
	var typWordInfos []def.TypWordInfo

	sql := GetUpdateTypWordInfoSQL()

	userId := values["userId"]
	temp := values["typWordInfo"]
 	json.Unmarshal([]byte(temp), &typWordInfos)

	//SQL実行
	for _, typWordInfo := range typWordInfos {
		_, err := db.Exec(sql, typWordInfo.SuccessTypCount, typWordInfo.MissTypCount, userId, typWordInfo.Word)
		if err != nil {
			if mysqlErr, ok := err.(*mysql.MySQLError); ok {
				logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
			}
			log.Fatal(err)
		}
	}

	logs.WriteLog("UpdateTypWordInfo正常終了", def.NORMAL)
	return
}

// アルファベットのタイピング情報を更新する関数
func UpdateTypAlphabetInfo(db *sql.DB, values map[string]string) {
	logs.WriteLog("UpdateTyoAlphabetInfo開始", def.NORMAL)
	var typAlphabetInfos []def.TypAlphabetInfo
	sql := GetUpdateTypAlphabetInfoSQL()

	userId := values["userId"]
	temp := values["typAlphabetInfo"]
	json.Unmarshal([]byte(temp), &typAlphabetInfos)

	for _, typAlphabetInfo := range typAlphabetInfos {
		// タイピング成功回数とタイピング失敗回数が0の場合は更新しない
		// どちらも0の場合は、タイピング情報がないということなので更新しない
		// Earlyreturn ミノコードの賜物
		if typAlphabetInfo.SuccessTypCount == 0 && typAlphabetInfo.MissTypCount == 0 { continue }
		
		_, err := db.Exec(sql, typAlphabetInfo.SuccessTypCount, typAlphabetInfo.MissTypCount, userId, typAlphabetInfo.Alphabet)
		if err != nil {
			if mysqlErr, ok := err.(*mysql.MySQLError); ok {
				logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
			}
			log.Fatal(err)
		}
	
	}

	logs.WriteLog("UpdateTyoAlphabetInfo正常終了", def.NORMAL)
	return
}

// ログイン情報を更新するSQLを返す関数
func GetUpdateLogoutDataSQL() string {
	return def.UPDATE_LOGOUT_DATA_SQL
}
// 単語のタイピング情報を更新するSQLを返す関数
func GetUpdateTypWordInfoSQL() string {
	return def.UPDATE_TYP_WORD_INFO_SQL
}

// アルファベットのタイピング情報を更新するSQLを返す関数
func GetUpdateTypAlphabetInfoSQL() string {
	return def.UPDATE_TYP_ALPHABET_INFO_SQL
}