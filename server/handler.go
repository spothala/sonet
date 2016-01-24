package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sonet/facebook"
	"sonet/instagram"
	"sonet/twitter"
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
				req.ParseForm()
				status := req.Form.Get("status")
				if status == "" {
					utils.RespondJson(w, ServerStatus{"Please Post along with status Parameter"})
					return
				}
				if facebook.CheckLoginStatus() && twitter.CheckLoginStatus() && instagram.CheckLoginStatus() {
					PostRequest(status)
				} else {
					fmt.Println("Redirecting to authfb")
				}
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
		case "authtwitter":
			switch req.Method {
			case "POST":
				twitter.SignIn(w, req)
			case "GET":
				utils.RespondJson(w, ServerStatus{"End Point Responds"})
			default:
				utils.RespondError(w, nil, http.StatusMethodNotAllowed)
			}
		case "authinsta":
			switch req.Method {
			case "POST":
				fmt.Println(string(instagram.Auth(w, req)))
			case "GET":
				utils.RespondJson(w, ServerStatus{"End Point Responds"})
			default:
				utils.RespondError(w, nil, http.StatusMethodNotAllowed)
			}
		case "index":
			render(w, "index.html")
		case "oauth2callback":
			fmt.Println(req.Method)
			if req.URL.Query().Get("oauth_token") != "" {
				twitter.AccessToken = req.URL.Query().Get("oauth_token")
				twitter.ReIssueAccessToken(req.URL.Query().Get("oauth_verifier"))
			}
			utils.RespondJson(w, ServerStatus{twitter.AccessToken})
		case "callback":
			fmt.Println(req.URL.Query().Get("code"))
			if req.URL.Query().Get("code") != "" {
				form := url.Values{}
				form.Set("client_id", instagram.Client_id)
				form.Set("client_secret", instagram.Client_secret)
				form.Set("grant_type", "authorization_code")
				form.Set("redirect_uri", instagram.Redirect_uri)
				form.Set("code", req.URL.Query().Get("code"))
				jsonOut := utils.GetJson(utils.ProcessFormRequest("POST", "", instagram.ApiUrl+"/oauth/access_token", form))
				utils.WriteJsonToFile(jsonOut, instagram.AccessTokenFile)
			}
		default:
			if req.URL.Query().Get("code") != "" {
				respJson := utils.GetJson(facebook.ConfirmIdentity(w, req, req.URL.Query().Get("code"))) //TODO: Browser Sign-Up
				facebook.AccessToken = respJson.(map[string]interface{})["access_token"].(string)
				utils.WriteToFile(facebook.AccessToken, facebook.AccessTokenFile)
				expires_in := respJson.(map[string]interface{})["expires_in"].(float64)
				fmt.Println(((expires_in / 60) / 60) / 24)
				respJson = utils.GetJson(facebook.GetMyDetails(w, req))
				fmt.Println(respJson.(map[string]interface{})["id"].(string))
				fmt.Println(respJson.(map[string]interface{})["name"].(string))
			}
			utils.RespondJson(w, ServerStatus{facebook.AccessToken})
		}
		return
	})
}
