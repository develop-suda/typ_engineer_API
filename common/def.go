package def

import ()

const (
	ERROR_LOGS_PATH  = "internal/log/error-logs/"
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
