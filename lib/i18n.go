package lib

var LANGUAGE = "jp"

var MessageJp = map[string]string {
	"SUCCESS_LOGIN": "ログインに成功しました。",
	"FAILED_LOGIN": "ログインに失敗しました。",
	"SUCCESS_LOGOUT": "ログアウトに成功しました。",
	"FAILED_LOGOUT": "ログアウトに失敗しました。",
}

func Message(key string) string {
	if LANGUAGE == "jp" {
		return MessageJp[key]
	} else {
		return "Un define message."
	}
}