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

	LoginData struct {
		User_id string `json:"user_id"`
		TokenString string `json:"tokenString"`
	}

	 WordDetail struct {
		Word string `json:"word"`
		Description string `json:"description"`
		Parts_of_speech string `json:"parts_of_speech"`
		Word_type string `json:"word_type"`
	}

	TypCount struct {
		Word string `json:"word"`
		Parts_of_speech string `json:"parts_of_speech"`
		Word_type string `json:"word_type"`
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
		Typing_count int `json:"typingCount"`
		Typing_miss_count int `json:"typingMissCount"`
	}

	AlphabetTypInfoSum struct {
		Typing_count int `json:"typingCount"`
		Typing_miss_count int `json:"typingMissCount"`
	}

	WordCountRanking struct {
		Word string `json:"word"`
		Typing_count int `json:"typingCount"`
		Rank_result int `json:"rankResult"`
	}

	WordMissCountRanking struct {
		Word string `json:"word"`
		Typing_miss_count int `json:"typingMissCount"`
		Rank_result int `json:"rankResult"`
	}

	AlphabetCountRanking struct {
		Alphabet string `json:"alphabet"`
		Typing_count int `json:"typingCount"`
		Rank_result int `json:"rankResult"`
	}

	AlphabetMissCountRanking struct {
		Alphabet string `json:"alphabet"`
		Typing_miss_count int `json:"typingMissCount"`
		Rank_result int `json:"rankResult"`
	}

)

// リクエストバラメータを受け取るための構造体
type(

	TypWordSelect struct {
		Word_type string
		Parts_of_speech string
		Alphabet string
		Quantity int
	}

	UserRegisterInfo struct {
		Last_name string
		First_name string
		Email string
		Password string
	}

	UserMatchInfo struct {
		Email string
		Password string
	}

	UserIdStruct struct {
		User_id string
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

)
