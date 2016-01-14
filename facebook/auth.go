package facebook

import (
    "sonet/utils"
    "net/http"
    "fmt"
    "net/url"
)

const (
    graphApiUrl  = "https://graph.facebook.com"
    redirect_uri = "http://localhost:8080/"
    client_id = "1071648322880677"
)

func Auth(client *http.Client, w http.ResponseWriter, req *http.Request) {
    fmt.Println("Trying to Auth Facebook")
    params := url.Values{}
    params.Set("type","user_agent")
    params.Set("client_id",client_id)
    params.Set("redirect_uri", redirect_uri)
    params.Set("scope","user_about_me,user_friends,user_likes,user_posts,user_events,user_status")
    httpReq, err := http.NewRequest("POST", graphApiUrl+"/oauth/authorize?"+params.Encode(), nil)
    if err != nil {
       fmt.Println("Failed to Prepare JsonRequest")
    }
    resp, err := client.Do(httpReq)
    if err != nil {
       fmt.Println(err)
    }
    fmt.Println("Status Code: "+resp.Status)
    w.Header().Set("Content-Type", "application/html")
    w.Write(utils.ReturnResponseBody(resp.Body))
}
