package logout

import(
	"log"
	"fmt"
	"database/sql"

	"github.com/go-sql-driver/mysql"

	def "github.com/develop-suda/typ_engineer_API/common"
	logs "github.com/develop-suda/typ_engineer_API/internal/log"
)

func UpdateLogoutData(db *sql.DB, userId def.UserIdStruct) error {
	logs.WriteLog("UpdateLogoutData開始", def.NORMAL)

	err := userId.Validate()
	if err != nil {
		log.Fatal(err)
		return err
	}

	sql := def.UPDATE_LOGOUT_DATA_SQL

	//トランザクション開始
	tx, err := db.Begin()
	defer tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	//SQL実行
	_, err = tx.Exec(sql, userId.User_id)
	
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
		return err
	}
	
	logs.WriteLog("UpdateLogoutData終了", def.NORMAL)
	return nil
}