package selectItems

import (
	"database/sql"
	"github.com/develop-suda/typ_engineer_API/common"
	"fmt"
	"log"
)

type user struct {
	Name         string
	Email        string
	Phone_number string
}

func GetTypWords(db *sql.DB, values map[string]string) []def.Word {

	// 複数件取得する場合、構造体を配列にする
	var words []def.Word
	var onlyWordTypeId []def.WordTypeId

	sql := def.GET_WORD_TYPE_ID_SQL
	sql = fmt.Sprintf(sql, values["type"])

	result, err := db.Query(sql)

	for result.Next() {
		wordTypeId := def.WordTypeId{}
		if err := result.Scan(&wordTypeId.Word_type_id); err != nil {
			log.Fatal(err)
		}
		onlyWordTypeId = append(onlyWordTypeId, wordTypeId)
	}

	if err != nil {
		log.Fatal(err)
	}

	sql = def.GET_TYP_WORDS_SQL
	sql = fmt.Sprintf(sql,
		values["type"],
		values["parts_of_speech"],
		values["alphabet"],
		values["quantity"])

	result, err = db.Query(sql)

	for result.Next() {
		word := def.Word{}
		if err := result.Scan(&word.Word, &word.Parts_of_speech, &word.Description); err != nil {
			log.Fatal(err)
		}
		words = append(words, word)
	}

	return words
}

func GetTypes(db *sql.DB) []def.WordType {

	// 複数件取得する場合、構造体を配列にする
	var wordTypes []def.WordType

	sql := "SELECT word_type FROM word_types ORDER BY word_type ASC"
	result, err := db.Query(sql)

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

	return wordTypes
}

func GetPartsOfSpeeches(db *sql.DB) []def.PartsOfSpeech {

	// 複数件取得する場合、構造体を配列にする
	var partsOfSpeeches []def.PartsOfSpeech

	sql := "SELECT parts_of_speech FROM parts_of_speeches ORDER BY parts_of_speech ASC"
	result, err := db.Query(sql)

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

	return partsOfSpeeches
}
