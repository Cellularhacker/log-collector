package v1

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log-collector/model/_util/pageInfo"
	"log-collector/model/systemLog"
	"log-collector/util"
	"log-collector/util/apiError"
	"net/http"
)

func SystemLogsGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) *apiError.Error {
	//userID := util.GetManagerID(r)
	//if userID == nil {
	//	return apiError.NotAuthorizedUser()
	//}

	//Getting PagingInfo...
	pi, iLoc := pageInfo.ParsePageInfo(r)
	if iLoc != nil {
		return apiError.BadRequestError(*iLoc)
	}

	sls, pRes, err := systemLog.GetByPageInfo(pi)
	if err != nil {
		return apiError.DetectError(err)
	}

	data := interface{}(sls)

	return util.SendDataResponse(w, &data, pRes, 200)
}

func SystemLogsPOST(w http.ResponseWriter, r *http.Request, _ httprouter.Params) *apiError.Error {
	defer r.Body.Close()

	log := systemLog.New()
	err := json.NewDecoder(r.Body).Decode(log)
	if err != nil {
		return apiError.BadRequestError("body")
	}

	err = systemLog.Create(log)
	if err != nil {
		return apiError.InternalServerError(err)
	}

	return util.SendDataResponse(w, nil, nil, http.StatusCreated)
}
