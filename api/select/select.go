package selectItems

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"sort"

	def "github.com/develop-suda/typ_engineer_API/common"
	logs "github.com/develop-suda/typ_engineer_API/internal/log"
)

type user struct {
	Name         string
	Email        string
	Phone_number string
}

func GetTypWords(db *sql.DB, values map[string]string) []def.Word {
	logs.WriteLog("GetTypWords開始", def.NORMAL)

	var words []def.Word // 複数件取得する場合、構造体を配列にする
	var keys []string    // 引数のキーを格納する配列

	// 引数のキーを配列に格納
	for key := range values {
		keys = append(keys, key)
	}
	// キーをソートする
	sort.Strings(keys)

	sql := def.GET_TYP_WORDS_SQL

	// 引数の値に合わせてSQLを変更
	// すべての場合は条件を追加する
	for _, key := range keys {
		if key == "1type" && values[key] != def.ALL { sql += " AND types.word_type = '" + values[key] + "'" }
		if key == "2parts_of_speech" && values[key] != def.ALL { sql += " AND pos.parts_of_speech = '" + values[key] + "'" }
		if key == "3alphabet" && values[key] != def.ALL { sql += " AND LEFT(words.word, 1) = '" + values[key] + "'" }
		if key == "4quantity" { sql += " ORDER BY RAND() LIMIT " + values[key] }
    }

	result, err := db.Query(sql)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	for result.Next() {
		word := def.Word{}
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
