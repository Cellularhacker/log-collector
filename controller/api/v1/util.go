package v1

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log-collector/model/_util/parseInfo"
	"log-collector/util"
	"log-collector/util/apiError"
	"net/http"
	"strings"
)

type ParseInfoReq struct {
	Type  string `json:"type"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

const (
	queryAllOK = "query_all_ok"
)

func ParseInfoPOST(w http.ResponseWriter, r *http.Request, _ httprouter.Params) *apiError.Error {
	userID := util.GetManagerID(r)
	if userID == nil {
		return apiError.NotAuthorizedUser()
	}

	defer r.Body.Close()

	body := &ParseInfoReq{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		return apiError.InternalServerError(err)
	}

	if body.Value == "" || body.Type == "" || body.Key == "" {
		return apiError.BadRequestError("value or type or key")
	}

	pi := parseInfo.New()
	pi.Type = body.Type
	pi.Key = body.Key
	pi.Value = body.Value
	pi.IsActive = true

	err = pi.Create()
	if err != nil {
		return apiError.InternalServerError(err)
	}

	return util.SendSuccessResponse(w, true)
}

func CoffeeGET(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) *apiError.Error {
	return util.SendMsgResponse(w, "I reject to boil your teapot.", 418)
}

func CreateTestPOST(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) *apiError.Error {
	return util.SendSuccessResponse(w, true)
}

func OKTestPOST(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) *apiError.Error {
	return util.SendSuccessResponse(w, false)
}

func getContentType(r *http.Request) string {
	return strings.ToLower(r.Header.Get("Content-Type"))
}

func PingGET(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) *apiError.Error {
	return util.SendJSONResponse(w, "pong!")
}
