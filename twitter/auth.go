package twitter

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sonet/utils"
	"strconv"
	"strings"
	"time"
)

const (
	ApiUrl        = "https://api.twitter.com"
	client_id     = "" // CONSUMER_KEY
	client_secret = "" // CONSUMER_SECRET
	redirect_uri  = "http://127.0.0.1:8080/oauth2callback"
)

var AccessTokenFile = utils.GetHomeDir() + "/.twitter_access_token"
var AccessToken string
var AccessTokenSecret string
var UserID string
var ScreenName string

func CheckLoginStatus() (status bool) {
	if _, err := os.Stat(AccessTokenFile); err == nil {
		data, e := utils.ReadBytesFromFile(AccessTokenFile)
		tokenJson := utils.GetJson(data)
		AccessToken = tokenJson.(map[string]interface{})["access_token"].(string)
		AccessTokenSecret = tokenJson.(map[string]interface{})["oauth_token_secret"].(string)
		UserID = tokenJson.(map[string]interface{})["user_id"].(string)
		ScreenName = tokenJson.(map[string]interface{})["screen_name"].(string)
		fmt.Println(AccessToken, AccessTokenSecret, UserID, ScreenName)
		if e != nil {
			return false
		}
	} else {
		return false
	}
	return true
}

func getOAuthNonce() string {
	out := make([]byte, 32)
	_, err := rand.Read(out)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(out)
}

func getEpochTime() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func PrepareOAuthHeaders(reqMethod string, endUrl string, params url.Values) (header string) {
	params.Set("oauth_signature", createSignature(reqMethod, endUrl, params.Encode()))
	//fmt.Println(params.Encode())
	header = params.Encode()
	header = strings.Replace(header, "&", "\", ", -1)
	header = strings.Replace(header, "=", "=\"", -1)
	//fmt.Println(header + "\"")
	return "OAuth " + header + "\""
}

func createSignature(reqMethod string, endUrl string, paramsEncode string) (sig string) {
	h := hmac.New(sha1.New, []byte(url.QueryEscape(client_secret)+"&"+url.QueryEscape(AccessTokenSecret))) //TODO: Little Suspicious
	h.Write([]byte(reqMethod + "&" + url.QueryEscape(ApiUrl+endUrl) + "&" + url.QueryEscape(paramsEncode)))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func SignIn(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Trying to Auth Twitter")
	// Preparing Params
	params := url.Values{}
	params.Set("oauth_callback", redirect_uri)
	params.Set("oauth_consumer_key", client_id)
	params.Set("oauth_nonce", getOAuthNonce())
	params.Set("oauth_signature_method", "HMAC-SHA1")
	params.Set("oauth_timestamp", getEpochTime())
	params.Set("oauth_version", "1.0")
	method := "POST"
	endUrl := "/oauth/request_token"
	header := PrepareOAuthHeaders(method, endUrl, params)
	//fmt.Println(header)
	oauth_token_response := string(utils.ProcessRequest(method, header, ApiUrl+endUrl))
	fmt.Println(oauth_token_response)
	if strings.Contains(oauth_token_response, "oauth_token") {
		AccessToken = strings.Split(strings.Split(oauth_token_response, "&")[0], "=")[1]
		fmt.Println(AccessToken)
		params := url.Values{}
		params.Set("oauth_token", AccessToken)
		fmt.Println(string(utils.ProcessRequest("GET", "", ApiUrl+"/oauth/authenticate?"+params.Encode()))) // TODO: Replace this with LOGIN Page
	} else {
		fmt.Println(oauth_token_response)
	}
}

func ReIssueAccessToken(oauth_verifier string) {
	fmt.Println("Re-Issuing the Access Token ....")
	// Preparing Params
	params := GetHeadersMap()
	method := "POST"
	endUrl := "/oauth/access_token"
	header := PrepareOAuthHeaders(method, endUrl, params)
	fmt.Println(header)
	params = url.Values{}
	params.Set("oauth_verifier", oauth_verifier)
	oauth_token_response := string(utils.ProcessRequest(method, header, ApiUrl+endUrl+"?"+params.Encode()))
	if strings.Contains(oauth_token_response, "oauth_token") {
		AccessToken = strings.Split(strings.Split(oauth_token_response, "&")[0], "=")[1]
		AccessTokenSecret = strings.Split(strings.Split(oauth_token_response, "&")[1], "=")[1]
		UserID = strings.Split(strings.Split(oauth_token_response, "&")[2], "=")[1]
		ScreenName = strings.Split(strings.Split(oauth_token_response, "&")[3], "=")[1]
		// Write Everything to JSON
		jsonOut := map[string]string{"access_token": AccessToken, "oauth_token_secret": AccessTokenSecret, "user_id": UserID, "screen_name": ScreenName}
		utils.WriteJsonToFile(jsonOut, AccessTokenFile)
	}
}

func GetHeadersMap() url.Values {
	params := url.Values{}
	params.Set("oauth_consumer_key", client_id)
	params.Set("oauth_nonce", getOAuthNonce())
	params.Set("oauth_signature_method", "HMAC-SHA1")
	params.Set("oauth_timestamp", getEpochTime())
	params.Set("oauth_token", AccessToken)
	params.Set("oauth_version", "1.0")
	return params
}
