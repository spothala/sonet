package instagram

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sonet/utils"
)

const (
	ApiUrl        = "https://api.instagram.com"
	Redirect_uri  = "http://127.0.0.1:8080/callback"
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

func Auth(w http.ResponseWriter, req *http.Request) (response []byte) {
	fmt.Println("Trying to Auth Instagram")
	params := url.Values{}
	params.Set("client_id", Client_id)
	params.Set("redirect_uri", Redirect_uri)
	params.Set("response_type", "code")
	return utils.ProcessRequest("POST", "", ApiUrl+"/oauth/authorize?"+params.Encode(), nil)
}
