package util

import (
	"bytes"
	"io"
	"log-collector/util/apiError"
	"net/http"
	"strconv"
)

func SendHTMLResponse(w http.ResponseWriter, body []byte) *apiError.Error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))

	_, err := io.Copy(w, bytes.NewReader(body))
	if err != nil {
		return apiError.InternalServerError(err)
	}

	return nil
}
