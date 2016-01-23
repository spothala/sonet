package media

import "fmt"

type Instagram struct {
	Response string
}

func (r *Instagram) Post(status string) {
	fmt.Println("Post Added in Instagram Stream")
}

func init() {
	Add("instagram", func() SocialMedia {
		return &Instagram{}
	})
}
