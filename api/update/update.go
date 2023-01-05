package update

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"

	logs "github.com/develop-suda/typ_engineer_API/internal/log"
	def "github.com/develop-suda/typ_engineer_API/common"
)

// wordのタイピング情報を更新する関数
func UpdateTypWordInfo(db *sql.DB, values []def.TypWordInfo, userId def.UserIdStruct) error {
	logs.WriteLog("UpdateTypWordInfo開始", def.NORMAL)

	var typWordInfos []def.TypWordInfo
	var err error

	typWordInfos = values

	// バリデーションチェックをループで行う
	for _, typWordInfo := range typWordInfos {
		err = typWordInfo.Validate()
		if err != nil {
			logs.WriteLog(err.Error(), def.ERROR)
			return err
		}
	}

	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), def.ERROR)
		return err
	}

	sql := def.GetUpdateTypWordInfoSQL()

	//SQL実行
	for _, typWordInfo := range typWordInfos {
		_, err := db.Exec(sql, typWordInfo.SuccessTypCount, typWordInfo.MissTypCount, userId.User_id, typWordInfo.Word)
		if err != nil {
			if mysqlErr, ok := err.(*mysql.MySQLError); ok {
				logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
			}
			log.Fatal(err)
			return err
		}
	}

	logs.WriteLog("UpdateTypWordInfo正常終了", def.NORMAL)
	return nil
}

// アルファベットのタイピング情報を更新する関数
func UpdateTypAlphabetInfo(db *sql.DB, typAlphabetInfos []def.TypAlphabetInfo, userId def.UserIdStruct) error {
	logs.WriteLog("UpdateTyoAlphabetInfo開始", def.NORMAL)
	
	var err error

	// バリデーションチェックをループで行う
	for _, typAlphabetInfo := range typAlphabetInfos {
		err = typAlphabetInfo.Validate()
		if err != nil {
			logs.WriteLog(err.Error(), def.ERROR)
			return err
		}
	}

	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), def.ERROR)
		return err
	}

	sql := def.GetUpdateTypAlphabetInfoSQL()

	for _, typAlphabetInfo := range typAlphabetInfos {
		// タイピング成功回数とタイピング失敗回数が0の場合は更新しない
		// どちらも0の場合は、タイピング情報がないということなので更新しない
		// Earlyreturn ミノコードの賜物
		if typAlphabetInfo.SuccessTypCount == 0 && typAlphabetInfo.MissTypCount == 0 { continue }
		
		_, err := db.Exec(sql, typAlphabetInfo.SuccessTypCount, typAlphabetInfo.MissTypCount, userId.User_id, typAlphabetInfo.Alphabet)
		if err != nil {
			if mysqlErr, ok := err.(*mysql.MySQLError); ok {
				logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
			}
			log.Fatal(err)
			return err
		}
	
	}

	logs.WriteLog("UpdateTyoAlphabetInfo正常終了", def.NORMAL)
	return nil
}
