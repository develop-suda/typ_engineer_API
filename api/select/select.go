package selectItems

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/go-sql-driver/mysql"

	def "github.com/develop-suda/typ_engineer_API/common"
	logs "github.com/develop-suda/typ_engineer_API/internal/log"
)

// SQLは定数のままで、valuesの値の変換だけを行う
func GetTypWords(db *sql.DB, values def.TypWordSelect) ([]def.Word, error) {
	logs.WriteLog("GetTypWords開始", nil, def.NORMAL)

	var words []def.Word // 複数件取得する場合、構造体を配列にする
	var err error

	// バリデーションチェック
	err = values.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), values, def.ERROR)
		return words, err
	}

	// sqlを取得
	sql := def.GetTypWordsSQL()


	// valuesの値をSQLに反映
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

	// SQL実行
	result, err := db.Query(sql)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, values, def.ERROR)
		} else {
			logs.WriteLog(err.Error(), values, def.ERROR)
		}
		return words, err
	}

	// 取得したデータを構造体に格納
	for result.Next() {
		word := def.Word{}
		if err := result.Scan(&word.Word, &word.Parts_of_speech, &word.Description); err != nil {
			logs.WriteLog(err.Error(),
			def.Word{
				Word:           word.Word,
				Parts_of_speech: word.Parts_of_speech,
				Description:    word.Description,
			},
			def.ERROR)
			return words, err
		}
		words = append(words, word)
	}

	logs.WriteLog("GetTypWords正常終了", nil, def.NORMAL)
	return words, err
}

func GetTypes(db *sql.DB) ([]def.WordType, error) {
	logs.WriteLog("GetTypes開始", nil, def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var wordTypes []def.WordType

	sql := "SELECT word_type FROM word_types ORDER BY word_type ASC"

	// SQL実行
	result, err := db.Query(sql)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.NONE_SQL_ARGUMENT, def.ERROR)
		} else {
			logs.WriteLog(err.Error(), def.NONE_SQL_ARGUMENT, def.ERROR)
		}
		logs.WriteLog(err.Error(), def.NONE_SQL_ARGUMENT, def.ERROR)
		return wordTypes, err
	}

	// 取得したデータを構造体に格納
	for result.Next() {
		wordType := def.WordType{}
		if err := result.Scan(&wordType.Word_type); err != nil {
			logs.WriteLog(err.Error(),
			def.WordType{
				Word_type: wordType.Word_type,
			},
			def.ERROR)
			return wordTypes, err
		}
		wordTypes = append(wordTypes, wordType)	
	}

	logs.WriteLog("GetTypes正常終了", nil, def.NORMAL)
	return wordTypes, err
}

func GetPartsOfSpeeches(db *sql.DB) ([]def.PartsOfSpeech, error) {
	logs.WriteLog("GetPartsOfSpeeches開始", nil, def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var partsOfSpeeches []def.PartsOfSpeech
	var err error

	sql := "SELECT parts_of_speech FROM parts_of_speeches ORDER BY parts_of_speech ASC"

	// SQL実行
	result, err := db.Query(sql)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.NONE_SQL_ARGUMENT, def.ERROR)
		} else {
			logs.WriteLog(err.Error(), def.NONE_SQL_ARGUMENT, def.ERROR)
		}
		logs.WriteLog(err.Error(), def.NONE_SQL_ARGUMENT, def.ERROR)
		return partsOfSpeeches, err
	}

	// 取得したデータを構造体に格納
	for result.Next() {
		partsOfSpeech := def.PartsOfSpeech{}
		// Scanは読み取りね
		if err := result.Scan(&partsOfSpeech.Parts_of_speech); err != nil {
			logs.WriteLog(err.Error(),
			def.PartsOfSpeech{
				Parts_of_speech: partsOfSpeech.Parts_of_speech,
			},
			def.ERROR)
			return partsOfSpeeches, err
		}
		partsOfSpeeches = append(partsOfSpeeches, partsOfSpeech)
	}

	logs.WriteLog("GetPartsOfSpeeches正常終了", nil, def.NORMAL)
	return partsOfSpeeches, err
}

func MatchUserPassword(db *sql.DB, values def.UserMatchInfo) (string, error) {
	logs.WriteLog("MatchUserPassword開始", nil, def.NORMAL)

	var userId string
	var err error

	// バリデーションチェック
	err = values.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), values, def.ERROR)
		return userId, err
	}

	sql := "SELECT LPAD(user_id,8,0) FROM users WHERE email = ? AND password = ?"

	// SQL実行
	result := db.QueryRow(sql, values.Email, values.Password)
	if err = result.Err(); err != nil {
		logs.WriteLog(err.Error(), values, def.ERROR)
		return userId, err
	}
	
	err = result.Scan(&userId)
	if err != nil {
		logs.WriteLog(err.Error(), values, def.ERROR)
		return userId, err
	}


	logs.WriteLog("MatchUserPassword正常終了", nil, def.NORMAL)
	return userId, err
}

func GetWordDetail(db *sql.DB) ([]def.WordDetail, error) {
	logs.WriteLog("GetWordDetail開始", nil, def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var wordDetails []def.WordDetail
	var err error

	// SQL文を取得
	sql := def.GetWordDetailSQL()

	// SQL実行
	result, err := db.Query(sql)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, def.NONE_SQL_ARGUMENT, def.ERROR)
		} else {
			logs.WriteLog(err.Error(), def.NONE_SQL_ARGUMENT, def.ERROR)
		}
		logs.WriteLog(err.Error(), def.NONE_SQL_ARGUMENT, def.ERROR)
		return wordDetails, err
	}

	// 取得したデータを構造体に格納
	for result.Next() {
		wordDetail := def.WordDetail{}
		if err := result.Scan(&wordDetail.Word, &wordDetail.Description, &wordDetail.Parts_of_speech, &wordDetail.Word_type); err != nil {
			logs.WriteLog(err.Error(), 
				def.WordDetail{
					Word: wordDetail.Word,
					Description: wordDetail.Description,
					Parts_of_speech: wordDetail.Parts_of_speech,
					Word_type: wordDetail.Word_type,
				},
			def.ERROR)
			return wordDetails, err
		}
		wordDetails = append(wordDetails, wordDetail)
	}

	logs.WriteLog("GetWordDetail正常終了", nil, def.NORMAL)
	return wordDetails, err
}

func GetWordTypInfo(db *sql.DB, userId def.UserIdStruct) ([]def.TypCount, error) {
	logs.WriteLog("GetWordTypInfo開始", nil, def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var typWordInfos []def.TypCount
	var err error

	// バリデーションチェック
	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return typWordInfos, err
	}

	// SQL文を取得
	sql := def.GetWordTypInfoSQL()

	// SQL実行
	result, err := db.Query(sql,userId.User_id)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, userId, def.ERROR)
		} else {
			logs.WriteLog(err.Error(), userId, def.ERROR)
		}
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return typWordInfos, err
	}

	// 取得したデータを構造体に格納
	for result.Next() {
		typWordInfo := def.TypCount{}
		if err := result.Scan(&typWordInfo.SuccessTypCount, &typWordInfo.MissTypCount); err != nil {
			logs.WriteLog(err.Error(),
				def.TypCount{
					SuccessTypCount: typWordInfo.SuccessTypCount,
					MissTypCount: typWordInfo.MissTypCount,
				},
				def.ERROR)
			return typWordInfos, err
		}
		typWordInfos = append(typWordInfos, typWordInfo)
	}
	
	logs.WriteLog("GetWordTypsInfo正常終了", nil, def.NORMAL)
	return typWordInfos, err
}

func GetWordTypInfoSum(db *sql.DB, userId def.UserIdStruct) (def.WordTypInfoSum, error) {
	logs.WriteLog("GetWordTypInfoSum開始", nil, def.NORMAL)

	var wordTypInfoSum def.WordTypInfoSum
	var err error

	// バリデーションチェック
	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return wordTypInfoSum, err
	}

	// SQL文を取得
	sql := def.GetWordTypInfoSumSQL()

	// SQL実行
	result := db.QueryRow(sql,userId.User_id)
	if err = result.Err(); err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, userId, def.ERROR)
		} else {
			logs.WriteLog(err.Error(), userId, def.ERROR)
		}
		return wordTypInfoSum, err
	}

	// 取得したデータを構造体に格納
	if err := result.Scan(&wordTypInfoSum.Typing_count, &wordTypInfoSum.Typing_miss_count); err != nil {
		logs.WriteLog(err.Error(),
			def.WordTypInfoSum{
				Typing_count: wordTypInfoSum.Typing_count,
				Typing_miss_count: wordTypInfoSum.Typing_miss_count,
			},
			def.ERROR)
		return wordTypInfoSum, err
	}

	logs.WriteLog("GetWordTypInfoSum正常終了", nil, def.NORMAL)
	return wordTypInfoSum, err

}

func GetAlphabetTypInfoSum(db *sql.DB, userId def.UserIdStruct) (def.AlphabetTypInfoSum, error) {
	logs.WriteLog("GetAlphabetTypInfoSum開始", nil, def.NORMAL)

	var alphabetTypInfoSum def.AlphabetTypInfoSum
	var err error

	// バリデーションチェック
	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return alphabetTypInfoSum, err
	}

	// SQL文を取得
	sql := def.GetAlphabetTypInfoSumSQL()

	// SQL実行
	result := db.QueryRow(sql,userId.User_id)
	if err = result.Err(); err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, userId, def.ERROR)
		} else {
			logs.WriteLog(err.Error(), userId, def.ERROR)
		}
		return alphabetTypInfoSum, err
	}

	// 取得したデータを構造体に格納
	if err := result.Scan(&alphabetTypInfoSum.Typing_count, &alphabetTypInfoSum.Typing_miss_count); err != nil {
		logs.WriteLog(err.Error(),
			def.AlphabetTypInfoSum{
				Typing_count: alphabetTypInfoSum.Typing_count,
				Typing_miss_count: alphabetTypInfoSum.Typing_miss_count,
			},
			def.ERROR)
		return alphabetTypInfoSum, err
	}

	logs.WriteLog("GetAlphabetTypInfoSum正常終了", nil, def.NORMAL)
	return alphabetTypInfoSum, err
}

func GetWordCountRanking(db *sql.DB,userId def.UserIdStruct) ([]def.WordCountRanking, error) {
	logs.WriteLog("GetWordCountRanking開始", nil, def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var wordCountRankings []def.WordCountRanking
	var err error

	// バリデーションチェック
	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return wordCountRankings, err
	}

	// SQL文を取得
	sql := def.GetWordCountRankingSQL()

	result, err := db.Query(sql,userId.User_id)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, userId, def.ERROR)
		} else {
			logs.WriteLog(err.Error(), userId, def.ERROR)
		}
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return wordCountRankings, err
	}

	// 取得したデータを構造体に格納
	for result.Next() {
		wordCountRanking := def.WordCountRanking{}
		if err := result.Scan(&wordCountRanking.Word, &wordCountRanking.Typing_count, &wordCountRanking.Rank_result); err != nil {
			logs.WriteLog(err.Error(),
				def.WordCountRanking{
					Word: wordCountRanking.Word,
					Typing_count: wordCountRanking.Typing_count,
					Rank_result: wordCountRanking.Rank_result,
				},
				def.ERROR)
			return wordCountRankings, err
		}
		wordCountRankings = append(wordCountRankings, wordCountRanking)
	}

	logs.WriteLog("GetWordCountRanking正常終了", nil, def.NORMAL)
	return wordCountRankings, err
}

func GetWordMissCountRanking(db *sql.DB,userId def.UserIdStruct) ([]def.WordMissCountRanking, error) {
	logs.WriteLog("GetWordMissRanking開始", nil, def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var wordMissRankings []def.WordMissCountRanking
	var err error

	// バリデーションチェック
	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return wordMissRankings, err
	}

	// SQL文を取得
	sql := def.GetWordMissRankingSQL()

	// SQL実行
	result, err := db.Query(sql,userId.User_id)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, userId, def.ERROR)
		} else {
			logs.WriteLog(err.Error(), userId, def.ERROR)
		}
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return wordMissRankings, err
	}

	// 取得したデータを構造体に格納
	for result.Next() {
		wordMissRanking := def.WordMissCountRanking{}
		if err := result.Scan(&wordMissRanking.Word, &wordMissRanking.Typing_miss_count, &wordMissRanking.Rank_result); err != nil {
			logs.WriteLog(err.Error(),
				def.WordMissCountRanking{
					Word: wordMissRanking.Word,
					Typing_miss_count: wordMissRanking.Typing_miss_count,
					Rank_result: wordMissRanking.Rank_result,
				},
				def.ERROR)
			wordMissRankings = nil
			return wordMissRankings, err
		}
		wordMissRankings = append(wordMissRankings, wordMissRanking)
	}

	logs.WriteLog("GetWordMissRanking正常終了", nil, def.NORMAL)
	return wordMissRankings, err
}

func GetAlphabetCountRanking(db *sql.DB,userId def.UserIdStruct) ([]def.AlphabetCountRanking, error) {
	logs.WriteLog("GetAlphabetCountRanking開始", nil, def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var alphabetCountRankings []def.AlphabetCountRanking
	var err error

	// バリデーションチェック
	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return alphabetCountRankings, err
	}

	// SQL文を取得
	sql := def.GetAlphabetCountRankingSQL()

	// SQL実行
	result, err := db.Query(sql,userId.User_id)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, userId, def.ERROR)
		} else {
			logs.WriteLog(err.Error(), userId, def.ERROR)
		}
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return alphabetCountRankings, err
	}

	// 取得したデータを構造体に格納
	for result.Next() {
		alphabetCountRanking := def.AlphabetCountRanking{}
		if err := result.Scan(&alphabetCountRanking.Alphabet, &alphabetCountRanking.Typing_count, &alphabetCountRanking.Rank_result); err != nil {
			logs.WriteLog(err.Error(),
				def.AlphabetCountRanking{
					Alphabet: alphabetCountRanking.Alphabet,
					Typing_count: alphabetCountRanking.Typing_count,
					Rank_result: alphabetCountRanking.Rank_result,
				},
				def.ERROR)
			return alphabetCountRankings, err
		}
		alphabetCountRankings = append(alphabetCountRankings, alphabetCountRanking)
	}

	logs.WriteLog("GetAlphabetCountRanking正常終了", nil, def.NORMAL)
	return alphabetCountRankings, err
}

func GetAlphabetMissCountRanking(db *sql.DB,userId def.UserIdStruct) ([]def.AlphabetMissCountRanking, error) {
	logs.WriteLog("GetAlphabetMissCountRanking開始", nil, def.NORMAL)

	// 複数件取得する場合、構造体を配列にする
	var alphabetMissCountRankings []def.AlphabetMissCountRanking
	var err error

	// バリデーションチェック
	err = userId.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return alphabetMissCountRankings, err
	}
	
	// SQL文を取得
	sql := def.GetAlphabetMissRankingSQL()

	// SQL実行
	result, err := db.Query(sql,userId.User_id)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			logs.WriteLog(fmt.Sprintf("%d", mysqlErr.Number)+" "+mysqlErr.Message+"\n"+sql, userId, def.ERROR)
		} else {
			logs.WriteLog(err.Error(), userId, def.ERROR)
		}
		logs.WriteLog(err.Error(), userId, def.ERROR)
		return alphabetMissCountRankings, err
	}

	// 取得したデータを構造体に格納
	for result.Next() {
		alphabetMissCountRanking := def.AlphabetMissCountRanking{}
		if err := result.Scan(&alphabetMissCountRanking.Alphabet, &alphabetMissCountRanking.Typing_miss_count, &alphabetMissCountRanking.Rank_result); err != nil {
			logs.WriteLog(err.Error(),
				def.AlphabetMissCountRanking{
					Alphabet: alphabetMissCountRanking.Alphabet,
					Typing_miss_count: alphabetMissCountRanking.Typing_miss_count,
					Rank_result: alphabetMissCountRanking.Rank_result,
				},
				def.ERROR)
			alphabetMissCountRankings = nil
			return alphabetMissCountRankings, err
		}
		alphabetMissCountRankings = append(alphabetMissCountRankings, alphabetMissCountRanking)
	}

	logs.WriteLog("GetAlphabetMissCountRanking正常終了", nil, def.NORMAL)
	return alphabetMissCountRankings, err
}

func TranMatchUserPassword(tx *sql.Tx, values def.UserMatchInfo) (string, error) {
	logs.WriteLog("MatchUserPassword開始", nil, def.NORMAL)

	var userId string
	var err error

	// バリデーションチェック
	err = values.Validate()
	if err != nil {
		logs.WriteLog(err.Error(), values, def.ERROR)
		return userId, err
	}

	sql := "SELECT LPAD(user_id,8,0) FROM users WHERE email = ? AND password = ?"

	// SQL実行
	result := tx.QueryRow(sql, values.Email, values.Password)
	if err = result.Err(); err != nil {
		logs.WriteLog(err.Error(), values, def.ERROR)
		return userId, err
	}
	
	// 取得したデータを構造体に格納
	err = result.Scan(&userId)
	if err != nil {
		logs.WriteLog(err.Error(), values, def.ERROR)
		return userId, err
	}

	logs.WriteLog("MatchUserPassword正常終了", nil, def.NORMAL)
	return userId, err
}