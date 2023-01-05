package selectItems

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"strconv"

	def "github.com/develop-suda/typ_engineer_API/common"
	logs "github.com/develop-suda/typ_engineer_API/internal/log"
)


// SQLは定数のままで、valuesの値の変換だけを行う
func GetTypWords(db *sql.DB, values def.TypWordSelect) []def.Word {
	logs.WriteLog("GetTypWords開始", def.NORMAL)

	var words []def.Word // 複数件取得する場合、構造体を配列にする

	err := values.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), def.ERROR)
		return words
	}

	sql := def.GetTypWordsSQL()


	if values.Word_type != def.TYPE_ALL {
		sql += " AND types.word_type = '" + values.Word_type + "'"
	}
	if values.Parts_of_speech != def.TYPE_ALL {
		sql += " AND pos.parts_of_speech = '" + values.Parts_of_speech + "'"
	}
	if values.Alphabet != def.TYPE_ALL {
		sql += " AND LEFT(words.word, 1) = '" + values.Alphabet + "'"
	}


	sql += " ORDER BY RAND() LIMIT " + strconv.Itoa(values.Quantity)

	result, err := db.Query(sql)
	defer db.Close()
	if err != nil {
		// TODO 調べる
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
		return words
	}

	for result.Next() {
		word := def.Word{}
		// TODO 調べる
		if err := result.Scan(&word.Word, &word.Parts_of_speech, &word.Description); err != nil {
			log.Fatal(err)
			// ループの途中でエラーが発生するとwordsに値が入っている場合があるので、初期化する
			words = nil
			return words
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
		return wordTypes
	}

	for result.Next() {
		wordType := def.WordType{}
		if err := result.Scan(&wordType.Word_type); err != nil {
			log.Fatal(err)
			// ループの途中でエラーが発生するとwordTypesに値が入っている場合があるので、初期化する
			wordTypes = nil
			return wordTypes
		}
		wordTypes = append(wordTypes, wordType)
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
		return partsOfSpeeches
	}

	for result.Next() {
		partsOfSpeech := def.PartsOfSpeech{}
		// Scanは読み取りね
		if err := result.Scan(&partsOfSpeech.Parts_of_speech); err != nil {
			log.Fatal(err)
			// ループの途中でエラーが発生するとpartsOfSpeechesに値が入っている場合があるので、初期化する
			partsOfSpeeches = nil
			return partsOfSpeeches
		}
		partsOfSpeeches = append(partsOfSpeeches, partsOfSpeech)
	}

	logs.WriteLog("GetPartsOfSpeeches正常終了", def.NORMAL)
	return partsOfSpeeches
}

func MatchUserPassword(db *sql.DB, values def.UserMatchInfo) string {
	logs.WriteLog("MatchUserPassword開始", def.NORMAL)

	var userId string
	var err error

	err = values.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), def.ERROR)
		return userId
	}

	sql := "SELECT LPAD(user_id,8,0) FROM users WHERE email = ? AND password = ?"

	result := db.QueryRow(sql, values.Email, values.Password)
	if err = result.Err(); err != nil {
		fmt.Println(err)
		return userId
	}
	
	err = result.Scan(&userId)
	if err != nil {
		log.Fatal(err)
		return userId
	}


	logs.WriteLog("MatchUserPassword正常終了", def.NORMAL)
	return userId
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
		return wordDetails
	}

	for result.Next() {
		wordDetail := def.WordDetail{}
		if err := result.Scan(&wordDetail.Word, &wordDetail.Description, &wordDetail.Parts_of_speech, &wordDetail.Word_type); err != nil {
			log.Fatal(err)
			// ループの途中でエラーが発生するとwordDetailsに値が入っている場合があるので、初期化する
			wordDetails = nil
			return wordDetails
		}
		wordDetails = append(wordDetails, wordDetail)
	}

	logs.WriteLog("GetWordDetail正常終了", def.NORMAL)
	return wordDetails
}

func GetWordTypInfo(db *sql.DB, userId def.UserIdStruct) []def.TypCount {
	logs.WriteLog("GetWordTypInfo開始", def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var typWordInfos []def.TypCount
	var err error

	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), def.ERROR)
		return typWordInfos
	}

	sql := def.GetWordTypInfoSQL()

	result, err := db.Query(sql,userId.User_id)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
		return typWordInfos
	}

	for result.Next() {
		typWordInfo := def.TypCount{}
		if err := result.Scan(&typWordInfo.SuccessTypCount, &typWordInfo.MissTypCount); err != nil {
			log.Fatal(err)
			// ループの途中でエラーが発生するとtypWordInfosに値が入っている場合があるので、初期化する
			typWordInfos = nil
			return typWordInfos
		}
		typWordInfos = append(typWordInfos, typWordInfo)
	}
	
	logs.WriteLog("GetWordTypsInfo正常終了", def.NORMAL)
	return typWordInfos
}

func GetWordTypInfoSum(db *sql.DB, userId def.UserIdStruct) def.WordTypInfoSum {
	logs.WriteLog("GetWordTypInfoSum開始", def.NORMAL)

	var wordTypInfoSum def.WordTypInfoSum
	var err error

	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), def.ERROR)
		return wordTypInfoSum
	}

	sql := def.GetWordTypInfoSumSQL()

	result := db.QueryRow(sql,userId.User_id)

	if err := result.Scan(&wordTypInfoSum.Typing_count, &wordTypInfoSum.Typing_miss_count); err != nil {
		log.Fatal(err)
		return wordTypInfoSum
	}

	logs.WriteLog("GetWordTypInfoSum正常終了", def.NORMAL)
	return wordTypInfoSum

}

func GetAlphabetTypInfoSum(db *sql.DB, userId def.UserIdStruct) def.AlphabetTypInfoSum {
	logs.WriteLog("GetAlphabetTypInfoSum開始", def.NORMAL)

	var alphabetTypInfoSum def.AlphabetTypInfoSum
	var err error

	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), def.ERROR)
		return alphabetTypInfoSum
	}

	sql := def.GetAlphabetTypInfoSumSQL()

	result := db.QueryRow(sql,userId.User_id)
	if err := result.Scan(&alphabetTypInfoSum.Typing_count, &alphabetTypInfoSum.Typing_miss_count); err != nil {
		log.Fatal(err)
		return alphabetTypInfoSum
	}

	logs.WriteLog("GetAlphabetTypInfoSum正常終了", def.NORMAL)
	return alphabetTypInfoSum
}

func GetWordCountRanking(db *sql.DB,userId def.UserIdStruct) []def.WordCountRanking {
	logs.WriteLog("GetWordCountRanking開始", def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var wordCountRankings []def.WordCountRanking
	var err error

	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), def.ERROR)
		return wordCountRankings
	}

	sql := def.GetWordCountRankingSQL()

	result, err := db.Query(sql,userId.User_id)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
		return wordCountRankings
	}

	for result.Next() {
		wordCountRanking := def.WordCountRanking{}
		if err := result.Scan(&wordCountRanking.Word, &wordCountRanking.Typing_count, &wordCountRanking.Rank_result); err != nil {
			log.Fatal(err)
			// ループの途中でエラーが発生するとwordCountRankingsに値が入っている場合があるので、初期化する
			wordCountRankings = nil
			return wordCountRankings
		}
		wordCountRankings = append(wordCountRankings, wordCountRanking)
	}

	logs.WriteLog("GetWordCountRanking正常終了", def.NORMAL)
	return wordCountRankings
}

func GetWordMissCountRanking(db *sql.DB,userId def.UserIdStruct) []def.WordMissCountRanking {
	logs.WriteLog("GetWordMissRanking開始", def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var wordMissRankings []def.WordMissCountRanking
	var err error

	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), def.ERROR)
		return wordMissRankings
	}

	sql := def.GetWordMissRankingSQL()

	result, err := db.Query(sql,userId.User_id)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
		return wordMissRankings
	}

	for result.Next() {
		wordMissRanking := def.WordMissCountRanking{}
		if err := result.Scan(&wordMissRanking.Word, &wordMissRanking.Typing_miss_count, &wordMissRanking.Rank_result); err != nil {
			log.Fatal(err)
			// ループの途中でエラーが発生するとwordMissRankingsに値が入っている場合があるので、初期化する
			wordMissRankings = nil
			return wordMissRankings
		}
		wordMissRankings = append(wordMissRankings, wordMissRanking)
	}

	logs.WriteLog("GetWordMissRanking正常終了", def.NORMAL)
	return wordMissRankings
}

func GetAlphabetCountRanking(db *sql.DB,userId def.UserIdStruct) []def.AlphabetCountRanking {
	logs.WriteLog("GetAlphabetCountRanking開始", def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var alphabetCountRankings []def.AlphabetCountRanking
	var err error

	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), def.ERROR)
		return alphabetCountRankings
	}

	sql := def.GetAlphabetCountRankingSQL()

	result, err := db.Query(sql,userId)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
		return alphabetCountRankings
	}

	for result.Next() {
		alphabetCountRanking := def.AlphabetCountRanking{}
		if err := result.Scan(&alphabetCountRanking.Alphabet, &alphabetCountRanking.Typing_count, &alphabetCountRanking.Rank_result); err != nil {
			log.Fatal(err)
			// ループの途中でエラーが発生するとalphabetCountRankingsに値が入っている場合があるので、初期化する
			alphabetCountRankings = nil
			return alphabetCountRankings
		}
		alphabetCountRankings = append(alphabetCountRankings, alphabetCountRanking)
	}

	logs.WriteLog("GetAlphabetCountRanking正常終了", def.NORMAL)
	return alphabetCountRankings
}

func GetAlphabetMissCountRanking(db *sql.DB,userId def.UserIdStruct) []def.AlphabetMissCountRanking {
	logs.WriteLog("GetAlphabetMissCountRanking開始", def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var alphabetMissCountRankings []def.AlphabetMissCountRanking
	var err error

	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), def.ERROR)
		return alphabetMissCountRankings
	}
	
	sql := def.GetAlphabetMissRankingSQL()

	result, err := db.Query(sql,userId)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.ERROR)
		}
		log.Fatal(err)
		return alphabetMissCountRankings
	}

	for result.Next() {
		alphabetMissCountRanking := def.AlphabetMissCountRanking{}
		if err := result.Scan(&alphabetMissCountRanking.Alphabet, &alphabetMissCountRanking.Typing_miss_count, &alphabetMissCountRanking.Rank_result); err != nil {
			log.Fatal(err)
			// ループの途中でエラーが発生するとalphabetMissCountRankingsに値が入っている場合があるので、初期化する
			alphabetMissCountRankings = nil
			return alphabetMissCountRankings
		}
		alphabetMissCountRankings = append(alphabetMissCountRankings, alphabetMissCountRanking)
	}

	logs.WriteLog("GetAlphabetMissCountRanking正常終了", def.NORMAL)
	return alphabetMissCountRankings
}