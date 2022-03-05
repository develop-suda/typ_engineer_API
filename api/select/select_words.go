package selectItems

import (
	"log"
    "github.com/jinzhu/gorm"
    "net/url"
)

type word struct {
    Word    string  `json:"word"`
    Parts_of_speech string `json:"parts_of_speech"`
    Discription string `json:"Discription"`
}

type level struct {
    Level string `json:"level"`
}

func DbSelect(db *gorm.DB, urlParams string) *gorm.DB {

    // 複数件取得する場合、構造体を配列にする
    var words []word

    sql := "SELECT word, parts_of_speech, discription FROM words WHERE 1 = 1"
    u, err := url.Parse(urlParams)
    if err != nil {
        log.Fatal(err)
    }

    for key, values := range u.Query() {
        if key != "limit" {
            sql += " AND " + key + " = '" + values[0] + "'"
        } else {
            sql += " LIMIT " + values[0]
        }
    }

    result := db.Raw(sql).Find(&words)
    return result
}

func GetLevels(db *gorm.DB, urlParams string) *gorm.DB {

    // 複数件取得する場合、構造体を配列にする
    var levels []level

    sql := "SELECT level FROM words GROUP BY level"

    result := db.Raw(sql).Find(&levels)
    return result
}