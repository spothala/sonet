package server

import (
    "net/http"
    "encoding/json"
    "fmt"
    "errors"
    "sonet/utils"
)

type ServerStatus struct {
    Response string
}

type ServerVersion struct {
    Version string
    Description string
    Environment string
}

func Handler(config utils.Config) http.Handler {
    mux := http.NewServeMux()
    handler := handleJenkinsCall(mux, config)
    return handler;
}

func handleJenkinsCall(h http.Handler, config utils.Config) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
          fmt.Println("URL: "+req.URL.Path[1:])
          switch req.URL.Path[1:] {
            case "status":
                RespondJson(w, ServerStatus{"OK"})
            case "version":
                RespondJson(w, ServerVersion{config.Version, config.Description, config.Environment})
            case "post":
                switch req.Method {
                case "POST":
                    ParseRequest(w, req)
                case "GET":
                    RespondJson(w, ServerStatus{"End Point Responds"})
                default:
                    RespondError(w, nil, http.StatusMethodNotAllowed)
                }
            default:
              RespondError(w, errors.New("This routing is not defined"), http.StatusNotFound)
          }
          return
    })
}

func RespondError(w http.ResponseWriter, err error, status int) {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
}

func RespondJson(w http.ResponseWriter, JsonType interface{}) {
    js, err := json.Marshal(JsonType)
    if err != nil {
      RespondError(w, err, http.StatusInternalServerError)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}
