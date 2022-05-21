package insert

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"

	def "github.com/develop-suda/typ_engineer_API/common"
	logs "github.com/develop-suda/typ_engineer_API/internal/log"
)

func CreateUser(db *sql.DB, values map[string]string) {
	logs.WriteLog("CreateUser開始", def.NORMAL)
	sql := def.INSERT_USER_SQL

	sql = fmt.Sprintf(sql,
		values["first_name"],
		values["last_name"],
		values["email"],
		values["password"],
	)

	_, err := db.Query(sql)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	logs.WriteLog("CreateUser正常終了", def.NORMAL)
	return
}
