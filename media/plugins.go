package media

import (
    "net/http"
    "fmt"
)

type SocialMedia interface {
    Post(http.ResponseWriter, *http.Request)
}

var RegisteredPlatform = map[string]MediaCreator{}

type MediaCreator func() SocialMedia

/* Register the GitHub Event */
func Add(name string, media_creator MediaCreator) {
    RegisteredPlatform[name] = media_creator
    fmt.Println("Social Media ",name," registered")
}
