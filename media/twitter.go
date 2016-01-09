package media

import (
  "net/http"
  "fmt"
)

type Twitter struct {
    Response string
}

func (r *Twitter) Post(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Post Added in Twitter Stream")
}

func init() {
    Add("twitter", func() SocialMedia {
        return &Twitter{}
    })
}
