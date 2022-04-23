package json

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/jinzhu/gorm"
)

//jsonにエンコード
func Conversion(result *gorm.DB, buf *bytes.Buffer) *json.Encoder {

	enc := json.NewEncoder(buf)
	if err := enc.Encode(&result.Value); err != nil {
		log.Fatal(err)
	}

	return enc
}
