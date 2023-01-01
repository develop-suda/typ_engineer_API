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

	sql := def.GetTypWordsSQL()

	for _, key := range keys {
		if key == "type" && values[key] != def.TYPE_ALL { 
			sql += " AND types.word_type = '" + values[key] + "'"
		}
		if key == "parts_of_speech" && values[key] != def.TYPE_ALL {
			sql += " AND pos.parts_of_speech = '" + values[key] + "'"
		}
		if key == "alphabet" && values[key] != def.TYPE_ALL { 
			sql += " AND LEFT(words.word, 1) = '" + values[key] + "'"
		}
	}

	sql += " ORDER BY RAND() LIMIT " + values["quantity"]

	result, err := db.Query(sql)
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

	sql := def.GetTypWordsSQL()

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

func GetWordDetail(db *sql.DB) []def.WordDetail {
	
	logs.WriteLog("GetWordDetail開始", def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var wordDetails []def.WordDetail

	sql := def.GetWordDetailSQL()

	result, err := db.Query(sql)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	for result.Next() {
		wordDetail := def.WordDetail{}
		if err := result.Scan(&wordDetail.Word, &wordDetail.Description, &wordDetail.Parts_of_speech, &wordDetail.Word_type); err != nil {
			log.Fatal(err)
		}
		wordDetails = append(wordDetails, wordDetail)
	}

	logs.WriteLog("GetWordDetail正常終了", def.NORMAL)
	return wordDetails
}

func GetWordTypInfo(db *sql.DB, userId string) []def.TypCount {
	
	logs.WriteLog("GetWordTypInfo開始", def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var typWordInfos []def.TypCount

	sql := def.GetWordTypInfoSQL()

	result, err := db.Query(sql,userId)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	for result.Next() {
		typWordInfo := def.TypCount{}
		if err := result.Scan(&typWordInfo.SuccessTypCount, &typWordInfo.MissTypCount); err != nil {
			log.Fatal(err)
		}
		typWordInfos = append(typWordInfos, typWordInfo)
	}
	
	logs.WriteLog("GetWordTypsInfo正常終了", def.NORMAL)
	return typWordInfos
}

func GetWordTypInfoSum(db *sql.DB, userId string) def.WordTypInfoSum {
	
	logs.WriteLog("GetWordTypInfoSum開始", def.NORMAL)

	var wordTypInfoSum def.WordTypInfoSum

	sql := def.GetWordTypInfoSumSQL()

	result := db.QueryRow(sql,userId)
	fmt.Println(result)
	fmt.Println(sql)
	fmt.Println(userId)

	if err := result.Scan(&wordTypInfoSum.Typing_count, &wordTypInfoSum.Typing_miss_count); err != nil {
		log.Fatal(err)
	}

	logs.WriteLog("GetWordTypInfoSum正常終了", def.NORMAL)
	return wordTypInfoSum

}

func GetAlphabetTypInfoSum(db *sql.DB, userId string) def.AlphabetTypInfoSum {
	
	logs.WriteLog("GetAlphabetTypInfoSum開始", def.NORMAL)

	var alphabetTypInfoSum def.AlphabetTypInfoSum

	sql := def.GetAlphabetTypInfoSumSQL()

	result := db.QueryRow(sql,userId)
	if err := result.Scan(&alphabetTypInfoSum.Typing_count, &alphabetTypInfoSum.Typing_miss_count); err != nil {
		log.Fatal(err)
	}

	logs.WriteLog("GetAlphabetTypInfoSum正常終了", def.NORMAL)
	return alphabetTypInfoSum
}

func GetWordCountRanking(db *sql.DB,userId string) []def.WordCountRanking {
	
	logs.WriteLog("GetWordCountRanking開始", def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var wordCountRankings []def.WordCountRanking

	sql := def.GetWordCountRankingSQL()

	result, err := db.Query(sql,userId)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	for result.Next() {
		wordCountRanking := def.WordCountRanking{}
		if err := result.Scan(&wordCountRanking.Word, &wordCountRanking.Typing_count, &wordCountRanking.Rank_result); err != nil {
			log.Fatal(err)
		}
		wordCountRankings = append(wordCountRankings, wordCountRanking)
	}

	logs.WriteLog("GetWordCountRanking正常終了", def.NORMAL)
	return wordCountRankings
}

func GetWordMissCountRanking(db *sql.DB,userId string) []def.WordMissCountRanking {
	
	logs.WriteLog("GetWordMissRanking開始", def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var wordMissRankings []def.WordMissCountRanking

	sql := def.GetWordMissRankingSQL()

	result, err := db.Query(sql,userId)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	for result.Next() {
		wordMissRanking := def.WordMissCountRanking{}
		if err := result.Scan(&wordMissRanking.Word, &wordMissRanking.Typing_miss_count, &wordMissRanking.Rank_result); err != nil {
			log.Fatal(err)
		}
		wordMissRankings = append(wordMissRankings, wordMissRanking)
	}

	logs.WriteLog("GetWordMissRanking正常終了", def.NORMAL)
	return wordMissRankings
}

func GetAlphabetCountRanking(db *sql.DB,userId string) []def.AlphabetCountRanking {
	
	logs.WriteLog("GetAlphabetCountRanking開始", def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var alphabetCountRankings []def.AlphabetCountRanking

	sql := def.GetAlphabetCountRankingSQL()

	result, err := db.Query(sql,userId)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	for result.Next() {
		alphabetCountRanking := def.AlphabetCountRanking{}
		if err := result.Scan(&alphabetCountRanking.Alphabet, &alphabetCountRanking.Typing_count, &alphabetCountRanking.Rank_result); err != nil {
			log.Fatal(err)
		}
		alphabetCountRankings = append(alphabetCountRankings, alphabetCountRanking)
	}

	logs.WriteLog("GetAlphabetCountRanking正常終了", def.NORMAL)
	return alphabetCountRankings
}

func GetAlphabetMissCountRanking(db *sql.DB,userId string) []def.AlphabetMissCountRanking {
	
	logs.WriteLog("GetAlphabetMissCountRanking開始", def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var alphabetMissCountRankings []def.AlphabetMissCountRanking

	sql := def.GetAlphabetMissRankingSQL()

	result, err := db.Query(sql,userId)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
	}

	for result.Next() {
		alphabetMissCountRanking := def.AlphabetMissCountRanking{}
		if err := result.Scan(&alphabetMissCountRanking.Alphabet, &alphabetMissCountRanking.Typing_miss_count, &alphabetMissCountRanking.Rank_result); err != nil {
			log.Fatal(err)
		}
		alphabetMissCountRankings = append(alphabetMissCountRankings, alphabetMissCountRanking)
	}

	logs.WriteLog("GetAlphabetMissCountRanking正常終了", def.NORMAL)
	return alphabetMissCountRankings
}