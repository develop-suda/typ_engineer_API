package def

import ()

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

const (
	INT_CONST_VAL = 100
	STR_CONST_VAL = "Hello World as Constant"

	private_int_const_val    = 5
	private_string_const_val = "wow"
)
