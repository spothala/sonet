package media

import (
	"fmt"
	"net/url"
	"sonet/facebook"
	"sonet/utils"
)

type Facebook struct {
	Response string
}

func (r *Facebook) Post(status string) {
	fmt.Println("Post Added in Facebook Stream, " + facebook.AccessToken)
	params := url.Values{}
	params.Set("message", status)
	params.Set("access_token", facebook.AccessToken)
	response := utils.ProcessRequest("POST", "", facebook.GraphApiUrl+"/me/feed?"+params.Encode(), nil)
	fmt.Println(string(response))
}

func init() {
	Add("facebook", func() SocialMedia {
		return &Facebook{}
	})
}
