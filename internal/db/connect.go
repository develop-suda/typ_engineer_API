package connect

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"database/sql"
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
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(mysql_container:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := sql.Open(dbDriver, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	return db
}
