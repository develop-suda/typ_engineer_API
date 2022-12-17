package logout

import(
	"log"
	"fmt"
	"database/sql"

	"github.com/go-sql-driver/mysql"

	def "github.com/develop-suda/typ_engineer_API/common"
	logs "github.com/develop-suda/typ_engineer_API/internal/log"
)

func UpdateLogoutData(db *sql.DB, userId string) {
	logs.WriteLog("UpdateLogoutData開始", def.NORMAL)
	sql := def.UPDATE_LOGOUT_DATA_SQL

	//トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	//SQL実行
	_, err = tx.Exec(sql, userId)
	//commit
	defer tx.Commit()
	
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}
	
	logs.WriteLog("UpdateLogoutData終了", def.NORMAL)
}