package selectItems

import (
	"database/sql"
	// "github.com/develop-suda/typ_engineer_API/common"
	"fmt"
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type word struct {
	Word            string `json:"word"`
	Parts_of_speech string `json:"parts_of_speech"`
	Description     string `json:"Description"`
}

type wordType struct {
	Word_type string `json:"word_type"`
}

type partsOfSpeech struct {
	Parts_of_speech string `json:"parts_of_speech"`
}

type wordTypeId struct {
	Word_type_id string
}

type user struct {
	Name         string
	Email        string
	Phone_number string
}

func GetTypWords(db *sql.DB, values map[string]string) []word {

	// 複数件取得する場合、構造体を配列にする
	var words []word
	var onlyWordTypeId []wordTypeId

	// test := def.INT_CONST_VAL

	sql := GetWordTypeIdSQL()
	sql = fmt.Sprintf(sql, values["type"])

	result, err := db.Query(sql)

	for result.Next() {
        wordTypeId := wordTypeId{}
        if err := result.Scan(&wordTypeId.Word_type_id); err != nil {
            log.Fatal(err)
        }
        onlyWordTypeId = append(onlyWordTypeId,wordTypeId)
    }

	if err != nil {
        log.Fatal(err)
    }

	sql = GetTypWordsSQL()

	// for key, value := range values {
	// 	if key == "quantity" {
    // 		sql += fmt.Sprintf(" LIMIT '%s'", value)
	// 	} else {
	// 		sql += fmt.Sprintf(" AND %s = '%s'", key, value)
	// 	}
	// }

	sql = fmt.Sprintf(sql,
					   values["type"],
					   values["parts_of_speech"],
					   values["alphabet"],
					   values["quantity"])

	fmt.Println(sql)

	result, err = db.Query(sql)

	for result.Next() {
        word := word{}
        if err := result.Scan(&word.Word,&word.Parts_of_speech,&word.Description); err != nil {
            log.Fatal(err)
        }
        words = append(words,word)
    }

	return words
}

func GetTypes(db *sql.DB) []wordType {

	// 複数件取得する場合、構造体を配列にする
	var wordTypes []wordType

	sql := "SELECT word_type FROM word_types ORDER BY word_type ASC"
	result, err := db.Query(sql)

	for result.Next() {
        wordType := wordType{}
        if err := result.Scan(&wordType.Word_type); err != nil {
            log.Fatal(err)
        }
        wordTypes = append(wordTypes,wordType)
    }
	
    if err != nil {
        log.Fatal(err)
    }

	// fmt.Println(wordTypes)
	// fmt.Printf("%T\n", wordTypes)
	return wordTypes
}

func GetPartsOfSpeeches(db *sql.DB) []partsOfSpeech {

	// 複数件取得する場合、構造体を配列にする
	var partsOfSpeeches []partsOfSpeech

	sql := "SELECT parts_of_speech FROM parts_of_speeches ORDER BY parts_of_speech ASC"
	result, err := db.Query(sql)

	for result.Next() {
        partsOfSpeech := partsOfSpeech{}
        if err := result.Scan(&partsOfSpeech.Parts_of_speech); err != nil {
            log.Fatal(err)
        }
        partsOfSpeeches = append(partsOfSpeeches,partsOfSpeech)
    }
	
	if err != nil {
        log.Fatal(err)
    }

	return partsOfSpeeches
}
