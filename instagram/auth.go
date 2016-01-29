package instagram

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sonet/utils"
)

const (
	ApiUrl        = "https://api.instagram.com"
	Redirect_uri  = "http://localhost:8080/callback"
	Client_id     = "" //APP_ID
	Client_secret = "" //APP_SECRET
)

var AccessTokenFile = utils.GetHomeDir() + "/.insta_access_token"
var AccessToken string
var UserID string
var ScreenName string

func CheckLoginStatus() (status bool) {
	if _, err := os.Stat(AccessTokenFile); err == nil {
		data, e := utils.ReadBytesFromFile(AccessTokenFile)
		tokenJson := utils.GetJson(data)
		AccessToken = tokenJson.(map[string]interface{})["access_token"].(string)
		UserID = tokenJson.(map[string]interface{})["user"].(map[string]interface{})["id"].(string)
		ScreenName = tokenJson.(map[string]interface{})["user"].(map[string]interface{})["username"].(string)
		fmt.Println(AccessToken, UserID, ScreenName)
		if e != nil {
			return false
		}
	} else {
		return false
	}
	return true
}

func CreateSignature(params url.Values, endPoint string) string {
	sign_key := []byte(Client_secret)
	h := hmac.New(sha1.New, sign_key)
	sign := endPoint
	for k, v := range params {
		sign += "|" + k + "=" + v[0]
	}
	fmt.Println(sign)
	h.Write([]byte(sign))
	return base64.StdEncoding.EncodeToString(h.Sum(sign_key[:0]))
}

func Auth(w http.ResponseWriter, req *http.Request) (response []byte) {
	fmt.Println("Trying to Auth Instagram")
	params := url.Values{}
	params.Set("client_id", Client_id)
	params.Set("redirect_uri", Redirect_uri)
	params.Set("response_type", "code")
	return utils.ProcessRequest("POST", "", ApiUrl+"/oauth/authorize?"+params.Encode(), nil)
}
