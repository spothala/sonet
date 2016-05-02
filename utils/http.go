package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func NewClient() (client *http.Client) {
	// Creating HTTP Client with SSL support - Its Secure but we'll skip cert verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
	return
}

func ProcessFormRequest(method string, header string, Apiurl string, data url.Values) (response []byte) {
	httpReq, err := http.NewRequest(method, Apiurl, strings.NewReader(data.Encode()))
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if header != "" {
		httpReq.Header.Set("Authorization", header)
		//fmt.Println(header)
	}
	fmt.Println(httpReq.URL)
	if err != nil {
		fmt.Println("Failed to Prepare JsonRequest")
	}
	resp, err := NewClient().Do(httpReq)
	if err != nil {
		fmt.Println(err)
	}
	checkHttpResponseStatus(resp)
	return ReturnResponseBody(resp.Body)
}

func ProcessRequest(method string, header map[string][]string, Apiurl string, body io.ReadCloser) (response []byte) {
	httpReq, err := http.NewRequest(method, Apiurl, body)
	if header != nil {
		httpReq.Header = header
		fmt.Println(header)
	}
	fmt.Println(httpReq.URL)
	if err != nil {
		fmt.Println("Failed to Prepare JsonRequest")
	}
	resp, err := NewClient().Do(httpReq)
	if err != nil {
		fmt.Println(err)
	}
	checkHttpResponseStatus(resp)
	return ReturnResponseBody(resp.Body)
}

func checkHttpResponseStatus(httpResponse *http.Response) {
	//printHttpResponseBody(httpResponse)
	if httpResponse.StatusCode != 200 {
		fmt.Println("HTTP Status: " + httpResponse.Status)
		printHttpResponseBody(httpResponse)
		//os.Exit(1)
	}
}

func printHttpResponseBody(httpResponse *http.Response) {
	//defer httpResponse.Body.Close()
	contents, err := ioutil.ReadAll(httpResponse.Body)
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
