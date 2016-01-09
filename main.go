package main

import (
    "net/http"
    "flag"
    "fmt"
    "strconv"
    "sonet/server"
    "sonet/utils"
    _ "sonet/media"
)

func main() {
    filepath := flag.String("config", "./config.json", "config file path")
    port := flag.Int("port", 8080, "Server listening Port")
  	flag.Parse()
    config, err := utils.ConfigReader(*filepath)
    if err != nil {
        panic("Unable to open file")
    }
    fmt.Println("Server started listening on Port: "+strconv.Itoa(*port))
    http.ListenAndServe(":"+strconv.Itoa(*port), server.Handler(config));
}
