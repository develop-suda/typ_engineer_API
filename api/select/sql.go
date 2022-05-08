package selectItems

import()

func GetTypWordsSQL() string {
	SQL := 
   `SELECT  words.word,
			words.description,
			pos.parts_of_speech
	FROM words
	INNER JOIN parts_of_speeches pos
	ON pos.parts_of_speech_id = words.parts_of_speech_id
	INNER JOIN word_types types
	ON types.word_type_id = words.word_type_id
	WHERE 1 = 1
	AND types.word_type = '%s'
	AND pos.parts_of_speech = '%s'
	AND LEFT(words.word, 1) = '%s'
	ORDER BY RAND() LIMIT %s`

	return SQL
}

func GetWordTypeIdSQL() string {
	SQL := 
   `SELECT word_type_id
	FROM word_types
	WHERE word_type = '%s'`

	return SQL
}