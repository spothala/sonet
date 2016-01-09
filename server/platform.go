package server

import (
    "net/http"
    //"fmt"
    "sonet/media"
    //"errors"
)

func ParseRequest(w http.ResponseWriter, req *http.Request) {
    //fmt.Println("PLATFORM: ",req.URL.Query()["platform"][0])
    /*media_creator, ok := media.RegisteredPlatform[req.URL.Query()["platform"][0]]
    if !ok {
        RespondError(w, errors.New(req.URL.Query()["platform"][0]+" platform is not registered"), http.StatusInternalServerError)
    } else {
        social_media := media_creator()
        social_media.Post(w, req)
    }*/
    for _, media_creator := range media.RegisteredPlatform {
        //fmt.Println("PLATFORM: "+key)
        social_media := media_creator()
        social_media.Post(w, req)
    }
}
