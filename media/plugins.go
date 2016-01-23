package media

import "fmt"

type SocialMedia interface {
	//Auth(http.ResponseWriter, *http.Request)
	Post(string)
}

var RegisteredPlatform = map[string]MediaCreator{}

type MediaCreator func() SocialMedia

/* Register the GitHub Event */
func Add(name string, media_creator MediaCreator) {
	RegisteredPlatform[name] = media_creator
	fmt.Println("Social Media ", name, " registered")
}
