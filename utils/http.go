package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func NewClient() (client *http.Client) {
	// Creating HTTP Client with SSL support - Its Secure but we'll skip cert verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
	return
}

func ProcessHeaderRequest(method string, url string, header string) (response []byte) {
	httpReq, err := http.NewRequest(method, url, nil)
	httpReq.Header.Set("Authorization", header)
	fmt.Println(header)
	fmt.Println(httpReq.URL)
	if err != nil {
		fmt.Println("Failed to Prepare JsonRequest")
	}
	resp, err := NewClient().Do(httpReq)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Status Code: " + resp.Status)
	return ReturnResponseBody(resp.Body)
}

func ProcessRequest(method string, url string) (response []byte) {
	httpReq, err := http.NewRequest(method, url, nil)
	fmt.Println(httpReq.URL)
	if err != nil {
		fmt.Println("Failed to Prepare JsonRequest")
	}
	resp, err := NewClient().Do(httpReq)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Status Code: " + resp.Status)
	return ReturnResponseBody(resp.Body)
}

func printHttpResponseBody(body io.ReadCloser) {
	//defer body.Close()
	contents, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}

func ReturnResponseBody(body io.ReadCloser) (response []byte) {
	//defer body.Close()
	contents, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	return contents
}

func RespondError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), http.StatusNotFound)
	return
}
