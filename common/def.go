package def

import ()

const (
	// エラーログの出力先パス
	ERROR_LOGS_PATH  = "internal/log/error-logs/"
	// 通常ログの出力先パス
	NORMAL_LOGS_PATH = "internal/log/normal-logs/"

	ERROR  = "error"
	NORMAL = "normal"

	TYPE_ALL = "ALL"
)

// エラーメッセージを定義
const (
	// エラーログ出力時、SQLの引数がない場合に使用
	NONE_SQL_ARGUMENT = "SQLの引数なし"
)
