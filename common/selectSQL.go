package def

import ()

// FROM句は基本的にViewを使用する!!
// Viewには、is_deleted = falseのレコードしかありません

const (

	// GET_SELECT_WORD_SQL : タイピングする単語情報を取得するSQL
	GET_TYP_WORDS_SQL = `SELECT  words.word,
			 words.description,
			 pos.parts_of_speech
	FROM words
	INNER JOIN parts_of_speeches pos
	ON pos.parts_of_speech_id = words.parts_of_speech_id
	INNER JOIN word_types types
	ON types.word_type_id = words.word_type_id
	WHERE types.word_type = ?
	AND pos.parts_of_speech = ?
	AND LEFT(words.word, 1) = ?
	ORDER BY RAND()
	LIMIT ?`

	// GET_WORD_UNIQUE_SQL : 一意になるよう単語を取得するSQL
	GET_WORD_UNIQUE_SQL = `
	SELECT word
	FROM v_words
	GROUP BY word`
	
)
