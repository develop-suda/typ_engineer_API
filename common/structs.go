package def

import (
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

	 WordDetail struct {
		Word string `json:"word"`
		Description string `json:"description"`
		Parts_of_speech string `json:"parts_of_speech"`
		Word_type string `json:"word_type"`
	}

	TypCount struct {
		SuccessTypCount int `json:"successTypCount"`
		MissTypCount int `json:"missTypCount"`
	}

	MyPageData struct {
		WordTypInfoSum WordTypInfoSum `json:"wordTypInfoSum"`
		AlphabetTypInfoSum AlphabetTypInfoSum `json:"alphabetTypInfoSum"`
		WordCountRanking []WordCountRanking `json:"wordCountRanking"`
		WordMissCountRanking []WordMissCountRanking `json:"wordMissCountRanking"`
		AlphabetCountRanking []AlphabetCountRanking `json:"alphabetCountRanking"`
		AlphabetMissCountRanking []AlphabetMissCountRanking `json:"alphabetMissCountRanking"`
	}

	WordTypInfoSum struct {
		Typing_count int
		Typing_miss_count int
	}

	AlphabetTypInfoSum struct {
		Typing_count int
		Typing_miss_count int
	}

	WordCountRanking struct {
		Word string 
		Typing_count int
		Rank_result int
	}

	WordMissCountRanking struct {
		Word string
		Typing_miss_count int
		Rank_result int
	}

	AlphabetCountRanking struct {
		Alphabet string
		Typing_count int
		Rank_result int
	}

	AlphabetMissCountRanking struct {
		Alphabet string
		Typing_miss_count int
		Rank_result int
	}

)
