package def

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// バリデーション

// 実際にタイピングする単語を取得する際のパラメータ
func (w TypWordSelect) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.Word_type, validation.Required, validation.In(WORD_TYPES...)),
		validation.Field(&w.Parts_of_speech, validation.Required, validation.In(PARTS_OF_SPEECHS...)),
		validation.Field(&w.Alphabet, validation.Required, validation.In(ALPHABETS...) ),
		validation.Field(&w.Quantity, validation.Required),
	)
}

func (w UserRegisterInfo) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.Last_name, validation.Required),
		validation.Field(&w.First_name, validation.Required),
		validation.Field(&w.Email, validation.Required, is.Email),
		validation.Field(&w.Password, validation.Required),
	)
}

func (w UserMatchInfo) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.Email, validation.Required, is.Email),
		validation.Field(&w.Password, validation.Required),
	)
}

func (w UserIdStruct) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.User_id, validation.Required),
	)
}

func (w TypWordInfo) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.Word, validation.Required),
		validation.Field(&w.SuccessTypCount, validation.Min(1)),
		validation.Field(&w.MissTypCount, validation.Min(0)),
	)
}

func (w TypAlphabetInfo) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.Alphabet, validation.Required),
		validation.Field(&w.SuccessTypCount, validation.Min(1)),
		validation.Field(&w.MissTypCount, validation.Min(0)),
	)
}

// バリデーション用の構造体
// ここから下の変数定義綺麗じゃない気がする。。。
// なんか追加したらここに追加していく感じになる
var WORD_TYPES = []interface{}{
	"basic",
	"advance",
	"beginner",
	"abbreviation",
	"initials",
	"ALL",
}

// バリデーション用の構造体
// 品詞確認用
var PARTS_OF_SPEECHS = []interface{}{
	"なし",
	"名詞",
	"代名詞",
	"動詞",
	"形容詞",
	"形容動詞",
	"連対詞",
	"副詞",
	"接続詞",
	"感動詞",
	"助動詞",
	"助詞",
	"前置詞",
	"動詞／名詞",
	"動詞／形容詞",
	"名詞／形容詞",
	"ALL",
}

// バリデーション用の構造体
// アルファベット確認用
var ALPHABETS = []interface{}{
	"a","b","c","d","e","f","g","h","i","j","k","l","m","n",
	"o","p","q","r","s","t","u","v","w","x","y","z","ALL",
}