package lib

import (
	"gopkg.in/resty.v1"
)

func GetToken(memberId string, password string) string {
	var url = "https://secure.sakura.ad.jp/auth/login"
	//	TODO	:	HTTPエラー時の処理
	//	TODO	:	認証エラー時の処理
	resp, _ := resty.R().
		SetFormData(map[string]string{
			"memberLogin[membercd]": memberId,
			"memberLogin[password]": password,
		}).
		Post(url)
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "SIDv2" {
			return cookie.Value
		}
	}
	return ""
}

func LogIn(memberId string, password string) bool {
	if "" != GetToken(memberId, password) {
		//	TODO	:	任意のディレクトリに取得したtokenを保存する処理
		return true
	} else {
		return false
	}
}

func LogOut() bool {
	var url = "https://secure.sakura.ad.jp/vps/api/v6.5/auth/logout"
	resp, _ := resty.R().
		SetHeaders(map[string]string{
			"X-Requested-With": "XMLHttpRequest",
			"Host": "secure.sakura.ad.jp",
			"Origin": "https://secure.sakura.ad.jp",
			"Content-Type": "application/json",
			"Cookie": "SIDv2=e0b0880e-9859-49ee-a07c-3cd2c8d0e6d7",
		}).
		Post(url)
	if resp.StatusCode() == 200 {
		return true
	} else {
		return false
	}
}