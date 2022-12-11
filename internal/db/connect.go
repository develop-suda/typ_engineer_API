package connect

import (
	"fmt"
	"os"

	"database/sql"
	"github.com/joho/godotenv"
	"github.com/go-sql-driver/mysql"
)

func DbConnect() *sql.DB {

	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	// pathはmain.goから見て
	err := godotenv.Load(".env")

	//もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// db variable.
	dbDriver := "mysql"

	CONNECT := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "mysql_container:3306",
		DBName:               os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open(dbDriver, CONNECT.FormatDSN())

	if err != nil {
		panic(err.Error())
	}

	return db
}
