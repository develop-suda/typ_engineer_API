package def

import ()

// FROM句は基本的にViewを使用する!!
// Viewには、is_deleted = falseのレコードしかありません

const (

	// SELECT_SELECT_WORD_SQL : タイピングする単語情報を取得するSQL
	SELECT_TYP_WORDS_SQL = `SELECT  words.word,
			 words.description,
			 pos.parts_of_speech
	FROM v_words
	INNER JOIN v_parts_of_speeches pos
	ON pos.parts_of_speech_id = words.parts_of_speech_id
	INNER JOIN v_word_types types
	ON types.word_type_id = words.word_type_id
	WHERE types.word_type = ?
	AND pos.parts_of_speech = ?
	AND LEFT(words.word, 1) = ?
	ORDER BY RAND()
	LIMIT ?`

	// SELECT_WORD_UNIQUE_SQL : 一意になるよう単語を取得するSQL
	SELECT_WORD_UNIQUE_SQL = `
	SELECT word
	FROM v_words
	GROUP BY word`
	
)

func GetTypWordsSQL() string {
	return SELECT_TYP_WORDS_SQL
}

func GetWordUniqueSQL() string {
	return SELECT_WORD_UNIQUE_SQL
}

