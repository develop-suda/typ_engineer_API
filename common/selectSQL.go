package def

import ()

// FROM句は基本的にViewを使用する!!
// Viewには、is_deleted = falseのレコードしかありません
// SQLを取得場合定数からではなく関数から取得する
// SQLを取得する関数はファイルの最後に記載する

const (

	// SELECT_TYP_WORD_SQL : タイピングする単語情報を取得するSQL
	SELECT_TYP_WORDS_SQL = `SELECT  words.word,
			 words.description,
			 pos.parts_of_speech
	FROM v_words words
	INNER JOIN v_parts_of_speeches pos
	ON pos.parts_of_speech_id = words.parts_of_speech_id
	INNER JOIN v_word_types types
	ON types.word_type_id = words.word_type_id
	WHERE 1 = 1`

	// SELECT_WORD_UNIQUE_SQL : 一意になるよう単語を取得するSQL
	SELECT_WORD_UNIQUE_SQL = `
	SELECT word
	FROM v_words
	GROUP BY word`

	// SELECT_WORD_DETAIL_SQL : 単語の詳細情報を取得するSQL
	// 詳細とは単語名と説明文、品詞、単語の難易度
	SELECT_WORD_DETAIL_SQL =`
	SELECT
		words.word,
		words.description,
		pos.parts_of_speech,
		types.word_type
	FROM v_words words
	LEFT JOIN v_parts_of_speeches pos ON words.parts_of_speech_id = pos.parts_of_speech_id
	LEFT JOIN v_word_types types ON words.word_type_id = types.word_type_id
	ORDER BY words.word ASC`

	// SELECT_WORD_TYP_INFO_SQL : 単語のタイピング情報を取得するSQL
	SELECT_WORD_TYP_INFO_SQL =`
	SELECT
    	typing_count,
    	typing_miss_count
	FROM v_typing_word_informations
	WHERE user_id = ?
	ORDER BY word ASC`

	// SELECT_ALPHABET_TYP_INFO_SQL : アルファベットのタイピング情報を取得するSQL
	SELECT_ALPHABET_TYP_INFO_SQL = `
	SELECT
		alphabet,
		typing_count,
		typing_miss_count
	FROM v_typing_alphabet_informations
	WHERE user_id = ?
	ORDER BY alphabet ASC`

	// SELECT_WORD_TYP_INFO_SUM_SQL : 単語のタイピング情報の合計を取得するSQL
	SELECT_WORD_TYP_INFO_SUM_SQL = `
	SELECT
		SUM(typing_count) AS typing_count,
		SUM(typing_miss_count) AS typing_miss_count
	FROM v_typing_word_informations
	WHERE user_id = ?`

	// SELECT_ALPHABET_TYP_INFO_SUM_SQL : アルファベットのタイピング情報の合計を取得するSQL
	SELECT_ALPHABET_TYP_INFO_SUM_SQL = `
	SELECT
		SUM(typing_count) AS typing_count,
		SUM(typing_miss_count) AS typing_miss_count
	FROM v_typing_alphabet_informations
	WHERE user_id = ?`

	// SELECT_WORD_COUNT_RANKING_SQL : 単語の成功カウントランキングTOP10を取得するSQL
	SELECT_WORD_COUNT_RANKING_SQL = `
	SELECT ranked_word_count.*
	FROM(
		SELECT
			word,
			typing_count,
			RANK() OVER(ORDER BY typing_count DESC) AS rank_result 
		FROM v_typing_word_informations
		WHERE user_id = ?
		AND typing_count <> 0
	) ranked_word_count
	WHERE ranked_word_count.rank_result <= 10`

	// SELECT_WORD_MISS_RANKING_SQL : 単語の失敗カウントランキングTOP10を取得するSQL
	SELECT_WORD_MISS_RANKING_SQL = `
	SELECT ranked_word_miss_count.*
	FROM(
		SELECT
			word,
			typing_miss_count,
			RANK() OVER(ORDER BY typing_miss_count DESC) AS rank_result 
		FROM v_typing_word_informations
		WHERE user_id = ?
		AND typing_miss_count <> 0
	) ranked_word_miss_count
	WHERE ranked_word_miss_count.rank_result <= 10`
		
	// SELECT_ALPHABET_COUNT_RANKING_SQL : アルファベットの成功カウントランキングTOP10を取得するSQL
	SELECT_ALPHABET_COUNT_RANKING_SQL = `
	SELECT ranked_alp_count.*
	FROM(
		SELECT
			alphabet,
			typing_count,
			RANK() OVER(ORDER BY typing_count DESC) AS rank_result 
		FROM v_typing_alphabet_informations
		WHERE user_id = ?
		AND typing_count <> 0
	) ranked_alp_count
	WHERE ranked_alp_count.rank_result <= 10`

	// SELECT_ALPHABET_MISS_RANKING_SQL : アルファベットの失敗カウントランキングTOP10を取得するSQL
	SELECT_ALPHABET_MISS_RANKING_SQL = `
	SELECT ranked_alp_miss_count.*
	FROM(
		SELECT
			alphabet,
			typing_miss_count,
			RANK() OVER(ORDER BY typing_miss_count DESC) AS rank_result 
		FROM v_typing_alphabet_informations
		WHERE user_id = ?
		AND typing_miss_count <> 0
	) ranked_alp_miss_count
	WHERE ranked_alp_miss_count.rank_result <= 10`

)

// SQLを取得する場合は、ここから取得する
func GetTypWordsSQL() string {
	return SELECT_TYP_WORDS_SQL
}

func GetWordUniqueSQL() string {
	return SELECT_WORD_UNIQUE_SQL
}

func GetWordDetailSQL() string {
	return SELECT_WORD_DETAIL_SQL
}

func GetWordTypInfoSQL() string {
	return SELECT_WORD_TYP_INFO_SQL
}

func GetAlphabetTypInfoSQL() string {
	return SELECT_ALPHABET_TYP_INFO_SQL
}

func GetWordTypInfoSumSQL() string {
	return SELECT_WORD_TYP_INFO_SUM_SQL
}

func GetAlphabetTypInfoSumSQL() string {
	return SELECT_ALPHABET_TYP_INFO_SUM_SQL
}

func GetWordCountRankingSQL() string {
	return SELECT_WORD_COUNT_RANKING_SQL
}

func GetWordMissRankingSQL() string {
	return SELECT_WORD_MISS_RANKING_SQL
}

func GetAlphabetCountRankingSQL() string {
	return SELECT_ALPHABET_COUNT_RANKING_SQL
}

func GetAlphabetMissRankingSQL() string {
	return SELECT_ALPHABET_MISS_RANKING_SQL
}

