package media

import (
  "net/http"
  "fmt"
)

type Instagram struct {
    Response string
}

func (r *Instagram) Post(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Post Added in Instagram Stream")
}

func init() {
    Add("instagram", func() SocialMedia {
        return &Instagram{}
    })
}
