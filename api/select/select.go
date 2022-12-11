package selectItems

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"

	def "github.com/develop-suda/typ_engineer_API/common"
	logs "github.com/develop-suda/typ_engineer_API/internal/log"
)


// SQLは定数のままで、valuesの値の変換だけを行う
func GetTypWords(db *sql.DB, values map[string]string) []def.Word {
	logs.WriteLog("GetTypWords開始", def.NORMAL)

	var words []def.Word // 複数件取得する場合、構造体を配列にする
	var keys []string    // 引数のキーを格納する配列

	// 引数のキーを配列に格納
	for key := range values {
		keys = append(keys, key)
	}

	sql := def.GET_TYP_WORDS_SQL

	for _, key := range keys {
		if key == "type" && values[key] == def.TYPE_ALL { values[key] = "types.word_type"}
		if key == "parts_of_speech" && values[key] == def.TYPE_ALL { values[key] = "pos.parts_of_speech"}
		if key == "alphabet" && values[key] == def.TYPE_ALL { values[key] = "LEFT(words.word, 1)"}
	}

	result, err := db.Query(sql, values["type"], values["parts_of_speech"], values["alphabet"], values["quantity"])
	if err != nil {
		// TODO 調べる
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	for result.Next() {
		word := def.Word{}
		// TODO 調べる
		if err := result.Scan(&word.Word, &word.Parts_of_speech, &word.Description); err != nil {
			log.Fatal(err)
		}
		words = append(words, word)
	}

	logs.WriteLog("GetTypWords正常終了", def.NORMAL)
	return words
}

func GetTypes(db *sql.DB) []def.WordType {
	logs.WriteLog("GetTypes開始", def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var wordTypes []def.WordType

	sql := "SELECT word_type FROM word_types ORDER BY word_type ASC"
	result, err := db.Query(sql)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	for result.Next() {
		wordType := def.WordType{}
		if err := result.Scan(&wordType.Word_type); err != nil {
			log.Fatal(err)
		}
		wordTypes = append(wordTypes, wordType)
	}

	if err != nil {
		log.Fatal(err)
	}

	logs.WriteLog("GetTypes正常終了", def.NORMAL)
	return wordTypes
}

func GetPartsOfSpeeches(db *sql.DB) []def.PartsOfSpeech {
	logs.WriteLog("GetPartsOfSpeeches開始", def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var partsOfSpeeches []def.PartsOfSpeech

	sql := "SELECT parts_of_speech FROM parts_of_speeches ORDER BY parts_of_speech ASC"
	result, err := db.Query(sql)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	for result.Next() {
		partsOfSpeech := def.PartsOfSpeech{}
		// Scanは読み取りね
		if err := result.Scan(&partsOfSpeech.Parts_of_speech); err != nil {
			log.Fatal(err)
		}
		partsOfSpeeches = append(partsOfSpeeches, partsOfSpeech)
	}

	if err != nil {
		log.Fatal(err)
	}

	logs.WriteLog("GetPartsOfSpeeches正常終了", def.NORMAL)
	return partsOfSpeeches
}

func MatchUserPassword(db *sql.DB, values map[string]string) string {
	logs.WriteLog("MatchUserPassword開始", def.NORMAL)

	var user_id string
	var err error

	sql := "SELECT LPAD(user_id,8,0) FROM users WHERE email = ? AND password = ?"

	result := db.QueryRow(sql, values["email"], values["password"])
	if err = result.Err(); err != nil {
		fmt.Println(err)
	}
	
	err = result.Scan(&user_id)
	if err != nil {
		log.Fatal(err)
	}


	logs.WriteLog("MatchUserPassword正常終了", def.NORMAL)
	return user_id
}

func ReturngetTypWordsSQL(db *sql.DB, values map[string]string) string {

	var words []def.Word // 複数件取得する場合、構造体を配列にする
	var keys []string    // 引数のキーを格納する配列

	// 引数のキーを配列に格納
	for key := range values {
		keys = append(keys, key)
	}

	sql := def.GET_TYP_WORDS_SQL

	for _, key := range keys {
		if key == "type" && values[key] == def.TYPE_ALL { values[key] = "types.word_type"}
		if key == "parts_of_speech" && values[key] == def.TYPE_ALL { values[key] = "pos.parts_of_speech"}
		if key == "alphabet" && values[key] == def.TYPE_ALL { values[key] = "LEFT(words.word, 1)"}
	}

	result, err := db.Query(sql, values["type"], values["parts_of_speech"], values["alphabet"], values["quantity"])
	if err != nil {
		// TODO 調べる
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	for result.Next() {
		word := def.Word{}
		// TODO 調べる
		if err := result.Scan(&word.Word, &word.Parts_of_speech, &word.Description); err != nil {
			log.Fatal(err)
		}
		words = append(words, word)
	}

	return sql
}