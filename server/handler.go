package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sonet/facebook"
	"sonet/utils"
)

type ServerStatus struct {
	Response string
}

type ServerVersion struct {
	Version     string
	Description string
	Environment string
}

func render(w http.ResponseWriter, tmpl string) {
	tmpl = fmt.Sprintf("templates/%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, "")
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func Handler(config utils.Config) http.Handler {
	mux := http.NewServeMux()
	handler := handleJenkinsCall(mux, config)
	return handler
}

func handleJenkinsCall(h http.Handler, config utils.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("URL: " + req.URL.Path[1:])
		switch req.URL.Path[1:] {
		case "status":
			utils.RespondJson(w, ServerStatus{"OK"})
		case "version":
			utils.RespondJson(w, ServerVersion{config.Version, config.Description, config.Environment})
		case "post":
			switch req.Method {
			case "POST":
				PostRequest(w, req)
			case "GET":
				utils.RespondJson(w, ServerStatus{"End Point Responds"})
			default:
				utils.RespondError(w, nil, http.StatusMethodNotAllowed)
			}
		case "authfb":
			switch req.Method {
			case "POST":
				facebook.Auth(w, req)
			case "GET":
				utils.RespondJson(w, ServerStatus{"End Point Responds"})
			default:
				utils.RespondError(w, nil, http.StatusMethodNotAllowed)
			}
		case "index":
			render(w, "index.html")
		default:
			if req.URL.Query().Get("code") != "" {
				respJson := utils.GetJson(facebook.ConfirmIdentity(w, req, req.URL.Query().Get("code")))
				facebook.AccessToken = respJson.(map[string]interface{})["access_token"].(string)
				//fmt.Println(facebook.AccessToken)
				respJson = utils.GetJson(facebook.GetMyDetails(w, req))
				fmt.Println(respJson.(map[string]interface{})["id"].(string))
				fmt.Println(respJson.(map[string]interface{})["name"].(string))
			}
			utils.RespondJson(w, ServerStatus{facebook.AccessToken})
		}
		return
	})
}
