package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sonet/utils"
)

const (
	GraphApiUrl   = "https://graph.facebook.com"
	redirect_uri  = "http://localhost:8080/"
	client_id     = "" //APP_ID
	client_secret = "" //APP_SECRET
)

var AccessToken string
var AccessTokenFile = utils.GetHomeDir() + "/.fb_access_token"

func CheckLoginStatus() (status bool) {
	if _, err := os.Stat(AccessTokenFile); err == nil {
		AccessToken, err = utils.ReadFromFile(AccessTokenFile)
		if err != nil {
			return false
		}
	} else {
		return false
	}
	return true
}

func GetMyDetails(w http.ResponseWriter, req *http.Request) (response []byte) {
	fmt.Println("Getting My Details")
	params := url.Values{}
	params.Set("fields", "id,name")
	params.Set("access_token", AccessToken)
	return utils.ProcessRequest("GET", "", GraphApiUrl+"/me?"+params.Encode())
}

func Auth(w http.ResponseWriter, req *http.Request) (response []byte) {
	fmt.Println("Trying to Auth Facebook")
	params := url.Values{}
	params.Set("client_id", client_id)
	params.Set("redirect_uri", redirect_uri)
	params.Set("scope", "user_about_me,user_friends,user_likes,user_posts,user_events,user_status,publish_actions")
	params.Set("display", "page")
	return utils.ProcessRequest("POST", "", GraphApiUrl+"/oauth/authorize?"+params.Encode())
}

func ConfirmIdentity(w http.ResponseWriter, req *http.Request, code string) (response []byte) {
	fmt.Println("Confirming Identity with FB")
	params := url.Values{}
	params.Set("client_id", client_id)
	params.Set("redirect_uri", redirect_uri)
	params.Set("client_secret", client_secret)
	params.Set("code", code)
	return utils.ProcessRequest("GET", "", GraphApiUrl+"/v2.3/oauth/access_token?"+params.Encode())
}
