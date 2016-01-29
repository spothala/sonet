package media

import (
	"fmt"
	"net/url"
	"sonet/instagram"
	"sonet/utils"
)

type Instagram struct {
	Response string
}

func (r *Instagram) Post(status string) {
	fmt.Println("Post Added in Instagram Stream")
	endUrl := "/v1/users/self/follows"
	form := url.Values{}
	form.Set("access_token", instagram.AccessToken)
	//form.Set("count", "2")
	form.Set("sig", instagram.CreateSignature(form, endUrl))
	jsonOut := utils.GetJson(utils.ProcessRequest("GET", "", instagram.ApiUrl+endUrl+"?"+form.Encode(), nil))
	fmt.Println(jsonOut)
}

func init() {
	Add("instagram", func() SocialMedia {
		return &Instagram{}
	})
}
