package def

import ()

const (
	GET_TYP_WORDS_SQL = 
	`SELECT  words.word,
			 words.description,
			 pos.parts_of_speech
	FROM words
	INNER JOIN parts_of_speeches pos
	ON pos.parts_of_speech_id = words.parts_of_speech_id
	INNER JOIN word_types types
	ON types.word_type_id = words.word_type_id
	WHERE 1 = 1`

	GET_WORD_TYPE_ID_SQL = 
   `SELECT word_type_id
	FROM word_types
	WHERE word_type = '%s'
	ORDER BY`
)
