package util

import (
	"encoding/json"
	"fmt"
	"log"
	"log-collector/util/apiError"
	"log-collector/util/uLog"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type errorResp struct {
	Message string
	Success bool `json:"success"`
	Code    int  `json:"code"`
}

//APIHandler a string making it easy to handle errors
type APIHandler func(http.ResponseWriter, *http.Request, httprouter.Params) *apiError.Error

//Handle returns a httprouter handler
func (h APIHandler) Handle() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := h(w, r, p); err != nil {
			errMsg := fmt.Sprintf("%s : %s : %s", r.RequestURI, err.Error, err.Message)
			if err.Code >= 500 && err.Code < 600 {
				uLog.Error(errMsg)
			} else {
				log.Println(errMsg)
			}

			param := &errorResp{"", false, err.Code}
			if err.Code != apiError.UnknownServerError {
				param.Message = err.Message
			}

			w.Header().Set("Content-Type", "application/json")
			if strings.Contains(strings.ToLower(err.Message), "invalid") {
				w.WriteHeader(http.StatusBadRequest)
			} else if err.Code >= 300 && err.Code < 600 {
				w.WriteHeader(err.Code)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			_ = json.NewEncoder(w).Encode(param)
		}
	}
}

func API(handler APIHandler) httprouter.Handle {
	return handler.Handle()
}
