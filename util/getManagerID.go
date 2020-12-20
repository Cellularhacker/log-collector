package util

import (
	"net/http"
)

func GetManagerID(r *http.Request) *string {
	//userID := tokenAuth.GetUserID(r)
	//if manager.IsManager(userID.Hex()) {
	//	uID := userID.Hex()
	//	return &uID
	//}

	//uLog.Warn(fmt.Sprintf("(NOT Manager) path: %s, id: %s", r.URL.Path, userID))
	return nil
}
