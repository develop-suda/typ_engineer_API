package def

import (
	"database/sql"
	"fmt"
)

func structfunc() {
	fmt.Println("structfunc")
}

type (
	Word struct {
		Word            string `json:"word"`
		Parts_of_speech string `json:"parts_of_speech"`
		Description     string `json:"Description"`
	}
	WordType struct {
		Word_type string `json:"word_type"`
	}

	PartsOfSpeech struct {
		Parts_of_speech string `json:"parts_of_speech"`
	}

	LoginData struct {
		User_id string `json:"user_id"`
		TokenString string `json:"tokenString"`
	}

	TypWordInfo struct {
		Word string
		SuccessTypCount int
		MissTypCount int
	}

	TypAlphabetInfo struct {
		Alphabet string
		SuccessTypCount int
		MissTypCount int	
	}

	typinfo struct {
		SuccessTypCount int
		MissTypCount int
	}
	TxAdmin struct {
		*sql.DB
	}

	Service struct {
		Tx TxAdmin
	}
)
