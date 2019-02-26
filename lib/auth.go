package lib

import (
	"gopkg.in/resty.v1"
	"os/user"
	"log"
	"os"
	"io/ioutil"
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

func SaveToken(token string) bool {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
		return false
	}

	file, err := os.Create(usr.HomeDir + "/.sakura_vps_token")
	if err != nil {
		return false
	}
	defer file.Close()

	output := token
	file.Write(([]byte)(output))
	return true
}

func LoadToken() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}

	dat, err := ioutil.ReadFile(usr.HomeDir + "/.sakura_vps_token")
	if err != nil {
		log.Fatal( err )
	}

	return string(dat)
}


func ClearToken() bool {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
		return false
	}

	file, err := os.Create(usr.HomeDir + "/.sakura_vps_token")
	if err != nil {
		return false
	}
	defer file.Close()

	output := ""
	file.Write(([]byte)(output))
	return true
}

func LogIn(memberId string, password string) bool {
	token := GetToken(memberId, password)
	if "" != token {
		if SaveToken(token) {
			return true
		}
	}
	return false
}

func LogOut() bool {
	var url = "https://secure.sakura.ad.jp/vps/api/v6.5/auth/logout"
	var token = LoadToken()
	if token == "" {
		return false
	}

	resp, _ := resty.R().
		SetHeaders(map[string]string{
			"X-Requested-With": "XMLHttpRequest",
			"Host": "secure.sakura.ad.jp",
			"Origin": "https://secure.sakura.ad.jp",
			"Content-Type": "application/json",
			"Cookie": token,
		}).
		Post(url)

	if resp.StatusCode() == 200 {
		ClearToken()
		return true
	} else {
		return false
	}
}