package def

import ()

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

	WordTypeId struct {
		Word_type_id string
	}

	LoginData struct {
		User_id string `json:"user_id"`
		TokenString string `json:"tokenString"`
	}
)
