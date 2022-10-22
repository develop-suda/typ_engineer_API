package def

import ()

const (

	// INSERT_USER_SQL : ユーザー情報を登録するSQL
	INSERT_USER_SQL = `
	INSERT INTO users
	(first_name,
	last_name,
	email,
	password,
	created_at,
	updated_at,
	is_deleted
	) VALUES (
	?,
	?,
	?,
	?,
	cast(now() as datetime),
	cast(now() as datetime),
	false)`

	// INSERT_LOGIN_DATA_SQL : ログイン情報を登録するSQL
	INSERT_LOGIN_DATA_SQL = `
	INSERT INTO login_history
	VALUES (
	?,
	cast(now() as datetime),
	NULL,
	cast(now() as datetime),
	cast(now() as datetime),
	false)`

	// INSERT_TYP_WORD_INFO_SQL : 単語のタイピング情報を登録するSQL (ユーザ登録時のみ使用)
	INSERT_TYPING_WORD_INFORMATIONS_SQL = `
	INSERT INTO typing_word_informations
	VALUES `

	// INSERT_TYP_ALPHABET_INFO_SQL : アルファベットのタイピング情報を登録するSQL (ユーザ登録時のみ使用)
	INSERT_TYPING_ALPHABET_INFORMATIONS_SQL = `
	INSERT INTO typing_alphabet_informations
	VALUES `
)
