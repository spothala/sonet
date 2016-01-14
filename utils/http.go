package utils

import (
  "io"
  "io/ioutil"
  "fmt"
  "os"
)

func printHttpResponseBody(body io.ReadCloser){
    //defer body.Close()
    contents, err := ioutil.ReadAll(body)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(contents))
}

func ReturnResponseBody(body io.ReadCloser)(response []byte){
    //defer body.Close()
    contents, err := ioutil.ReadAll(body)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }
    return contents
}
