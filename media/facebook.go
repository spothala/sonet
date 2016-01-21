package media

import (
	"fmt"
	"net/http"
	"net/url"
	"sonet/facebook"
	"sonet/utils"
)

type Facebook struct {
	Response string
}

func (r *Facebook) Post(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Post Added in Facebook Stream, " + facebook.AccessToken)
	params := url.Values{}
	params.Set("message", "First Multi-Platform Status Posting.... :) ")
	params.Set("access_token", facebook.AccessToken)
	response := utils.ProcessRequest("POST", facebook.GraphApiUrl+"/me/feed?"+params.Encode())
	fmt.Println(string(response))
}

func init() {
	Add("facebook", func() SocialMedia {
		return &Facebook{}
	})
}
