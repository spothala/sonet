package twitter

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sonet/utils"
	"strconv"
	"strings"
	"time"
)

const (
	ApiUrl        = "https://api.twitter.com"
	client_id     = ""
	redirect_uri  = "http://127.0.0.1:8080/oauth2callback"
	client_secret = ""
)

func prepareOAuthHeaders(reqMethod string, endUrl string) (header string) {
	out := make([]byte, 32)
	_, err := rand.Read(out)
	if err != nil {
		panic(err)
	}

	params := url.Values{}
	params.Set("oauth_callback", redirect_uri)
	params.Set("oauth_consumer_key", client_id)
	params.Set("oauth_nonce", base64.URLEncoding.EncodeToString(out))
	params.Set("oauth_signature_method", "HMAC-SHA1")
	params.Set("oauth_timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	params.Set("oauth_version", "1.0")
	params.Set("oauth_signature", createSignature(reqMethod, endUrl, params.Encode()))
	//fmt.Println(params.Encode())
	header = params.Encode()
	header = strings.Replace(header, "&", "\", ", -1)
	header = strings.Replace(header, "=", "=\"", -1)
	//fmt.Println(header + "\"")
	return "OAuth " + header + "\""
}

func createSignature(reqMethod string, endUrl string, paramsEncode string) (sig string) {
	base := reqMethod + "&" + url.QueryEscape(ApiUrl+endUrl) + "&" + url.QueryEscape(paramsEncode)
	//fmt.Println(base)
	signing_key := url.QueryEscape(client_secret) + "&"
	//fmt.Println(signing_key)
	h := hmac.New(sha1.New, []byte(signing_key))
	h.Write([]byte(base))
	sig = base64.StdEncoding.EncodeToString(h.Sum(nil))
	//fmt.Println(sig)
	return
}

func SignIn(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Trying to Auth Twitter")
	method := "POST"
	endUrl := "/oauth/request_token"
	header := prepareOAuthHeaders(method, endUrl)
	fmt.Println(string(utils.ProcessHeaderRequest(method, ApiUrl+endUrl, header)))
}
