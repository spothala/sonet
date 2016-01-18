package server

import (
	"net/http"
	"sonet/media"
)

func PostRequest(w http.ResponseWriter, req *http.Request) {
	for _, media_creator := range media.RegisteredPlatform {
		//fmt.Println("PLATFORM: "+key)
		social_media := media_creator()
		social_media.Post(w, req)
	}
}
