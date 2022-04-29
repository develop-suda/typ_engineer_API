package selectItems

import (
	"log"
	"net/url"

	"github.com/jinzhu/gorm"
)

type word struct {
	Word            string `json:"word"`
	Parts_of_speech string `json:"parts_of_speech"`
	Description     string `json:"Description"`
}

type level struct {
	Level_id int `json:"level"`
}

type user struct {
	Name         string
	Email        string
	Phone_number string
}

func DbSelect(db *gorm.DB, urlParams string) *gorm.DB {

	// 複数件取得する場合、構造体を配列にする
	var words []word

	sql := "SELECT word, parts_of_speech, description FROM words WHERE 1 = 1"
	u, err := url.Parse(urlParams)
	if err != nil {
		log.Fatal(err)
	}

	for key, values := range u.Query() {
		if key != "limit" {
			sql += " AND " + key + " = '" + values[0] + "'"
		} else {
			sql += " ORDER BY RAND() LIMIT " + values[0]
		}
	}

	result := db.Raw(sql).Find(&words)
	return result
}

func GetLevels(db *gorm.DB, urlParams string) *gorm.DB {

	// 複数件取得する場合、構造体を配列にする
	var levels []level

	sql := "SELECT level_id FROM words GROUP BY level_id"

	result := db.Raw(sql).Find(&levels)
	return result
}

func GetUsers(db *gorm.DB) *gorm.DB {
	var users []user

	result := db.Find(&users)
	return result
}
