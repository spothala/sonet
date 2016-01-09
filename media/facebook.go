package media

import (
  "net/http"
  "fmt"
)

type Facebook struct {
    Response string
}

func (r *Facebook) Post(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Post Added in Facebook Stream")
}

func init() {
    Add("facebook", func() SocialMedia {
        return &Facebook{}
    })
}
