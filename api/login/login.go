package login

import (
	"fmt"
	"time"
	"database/sql"

	"github.com/go-sql-driver/mysql"

	def "github.com/develop-suda/typ_engineer_API/common"
	logs "github.com/develop-suda/typ_engineer_API/internal/log"
	jwt "github.com/dgrijalva/jwt-go"	
)

func CreateToken(userId string) (def.LoginData, error) {
	logs.WriteLog("CreateToken開始", nil, def.NORMAL)

	var loginData def.LoginData
	var err error

	//jwt認証をする
	// TODO jwtを調べる
	// Claimsオブジェクトの作成
	claims := jwt.MapClaims {
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenの署名
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
		return loginData, err
	}

	loginData.User_id = userId
	loginData.TokenString = tokenString

	logs.WriteLog("CreateToken正常終了", nil, def.NORMAL)
	return loginData, err
}


func InsertLoginData(tx *sql.DB, userId string) error {
	logs.WriteLog("InsertLoginData開始", nil, def.NORMAL)
	sql := def.INSERT_LOGIN_DATA_SQL

	//SQL実行
	_, err := tx.Exec(sql, userId)
	
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, userId, def.ERROR)
		}
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return err
	}

	logs.WriteLog("InsertLoginData正常終了", nil, def.NORMAL)
	return nil
}

func TranInsertLoginData(tx *sql.Tx, userId string) error {
	logs.WriteLog("InsertLoginData開始", nil, def.NORMAL)
	sql := def.INSERT_LOGIN_DATA_SQL

	//SQL実行
	_, err := tx.Exec(sql, userId)
	
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, userId, def.ERROR)
		}
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return err
	}

	logs.WriteLog("InsertLoginData正常終了", nil, def.NORMAL)
	return nil
}