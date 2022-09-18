package token

import (
	"fmt"
	"time"

	def "github.com/develop-suda/typ_engineer_API/common"
	jwt "github.com/dgrijalva/jwt-go"	
)

func CreateToken(userId string) def.LoginData {

	var loginData def.LoginData

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
	}

	loginData.User_id = userId
	loginData.TokenString = tokenString

	return loginData
}