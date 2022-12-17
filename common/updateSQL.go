package def 

import ()

// このファイルではUPDATE文のみを定義しています。
// 必ずupdate_atを更新するようにしてください。

const (

	// UPDATE_LOGIN_INFO_SQL ログアウト情報を更新するSQL
	UPDATE_LOGOUT_DATA_SQL = `
	UPDATE login_history
	SET logout_date = cast(now() as datetime),
		updated_at = cast(now() as datetime)
	WHERE login_date = (
		SELECT max_login_date
		FROM (
			SELECT MAX(login_date) AS max_login_date
			FROM login_history
			WHERE user_id = ?
			AND is_deleted = false
			GROUP BY user_id
		) max_date
	)`

	// UPDATE_TYP_WORD_INFO_SQL : 単語のタイピング情報を更新するSQL
	UPDATE_TYP_WORD_INFO_SQL = `
	UPDATE typing_word_informations
	SET typing_count = typing_count + ?,
	    typing_miss_count = typing_miss_count + ?,
		updated_at = cast(now() as datetime)
	WHERE user_id = ?
	AND word = ?
	AND is_deleted = false`

	// UPDATE_TYP_ALPHABET_INFO_SQL : アルファベットのタイピング情報を更新するSQL
	UPDATE_TYP_ALPHABET_INFO_SQL = `
	UPDATE typing_alphabet_informations
	SET typing_count = typing_count + ?,
		typing_miss_count = typing_miss_count + ?,
		updated_at = cast(now() as datetime)
	WHERE user_id = ?
	AND alphabet = ?
	AND is_deleted = false`
)

// ログイン情報を更新するSQLを返す関数
func GetUpdateLogoutDataSQL() string {
	return UPDATE_LOGOUT_DATA_SQL
}
// 単語のタイピング情報を更新するSQLを返す関数
func GetUpdateTypWordInfoSQL() string {
	return UPDATE_TYP_WORD_INFO_SQL
}

// アルファベットのタイピング情報を更新するSQLを返す関数
func GetUpdateTypAlphabetInfoSQL() string {
	return UPDATE_TYP_ALPHABET_INFO_SQL
}