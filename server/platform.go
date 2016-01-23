package server

import "sonet/media"

func PostRequest(status string) {
	for _, media_creator := range media.RegisteredPlatform {
		//fmt.Println("PLATFORM: "+key)
		social_media := media_creator()
		social_media.Post(status)
	}
}
