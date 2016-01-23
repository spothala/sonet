package media

import (
	"fmt"
	"net/url"
	"sonet/twitter"
	"sonet/utils"
)

type Twitter struct {
	Response string
}

func (r *Twitter) Post(status string) {
	//listTweets()
	postTweet(status)
}

func postTweet(status string) {
	fmt.Println("Post Added in Twitter Stream " + twitter.AccessToken)
	params := twitter.GetHeadersMap()
	params.Set("status", url.QueryEscape(status))
	method := "POST"
	endUrl := "/1.1/statuses/update.json"
	header := twitter.PrepareOAuthHeaders(method, endUrl, params)
	params = url.Values{}
	params.Set("status", url.QueryEscape(status))
	response := utils.ProcessRequest(method, header, twitter.ApiUrl+endUrl+"?"+params.Encode())
	fmt.Println(string(response))
}

func listTweets() {
	fmt.Println("List of Tweets " + twitter.AccessToken)
	params := twitter.GetHeadersMap()
	params.Set("screen_name", twitter.ScreenName)
	params.Set("count", "2")
	method := "GET"
	endUrl := "/1.1/statuses/user_timeline.json"
	header := twitter.PrepareOAuthHeaders(method, endUrl, params)
	params = url.Values{}
	params.Set("screen_name", twitter.ScreenName)
	params.Set("count", "2")
	response := utils.ProcessRequest(method, header, twitter.ApiUrl+endUrl+"?"+params.Encode())
	fmt.Println(string(response))
}

func init() {
	Add("twitter", func() SocialMedia {
		return &Twitter{}
	})
}
