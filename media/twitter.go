package media

import (
	"fmt"
	"net/http"
	"net/url"
	"sonet/twitter"
	"sonet/utils"
)

type Twitter struct {
	Response string
}

func (r *Twitter) Post(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Post Added in Twitter Stream " + twitter.AccessToken)
	params := twitter.GetHeadersMap()
	params.Set("status", "First Multi-Platform Status Posting.... :) ")
	method := "POST"
	endUrl := "/1.1/statuses/update.json"
	header := twitter.PrepareOAuthHeaders(method, endUrl, params)
	//fmt.Println(header)
	params = url.Values{}
	params.Set("status", "First Multi-Platform Status Posting.... :) ")
	params.Set("oauth_token", twitter.AccessToken)
	response := utils.ProcessHeaderRequest(method, twitter.ApiUrl+endUrl+"?"+params.Encode(), header)
	fmt.Println(string(response))
}

func init() {
	Add("twitter", func() SocialMedia {
		return &Twitter{}
	})
}
