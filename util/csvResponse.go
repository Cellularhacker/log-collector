package util

import (
	"bytes"
	"fmt"
	"io"
	"log-collector/util/apiError"
	"log-collector/util/uTime"
	"net/http"
	"strconv"
)

func SendCSVResponseFromString(w http.ResponseWriter, fileNamePrefix string, csvStr *string) *apiError.Error {
	br := bytes.NewBufferString(*csvStr)
	fileName := fmt.Sprintf("%s_%s.csv", fileNamePrefix, uTime.GetKSTDateStr(nil))

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Length", strconv.Itoa(len(*csvStr)))

	_, err := io.Copy(w, br)
	if err != nil {
		return apiError.InternalServerError(err)
	}

	return nil
}

func SendCSVResponse(w http.ResponseWriter, fileNamePrefix string, csv *bytes.Buffer) *apiError.Error {
	fileName := fmt.Sprintf("%s_%s.csv", fileNamePrefix, uTime.GetKSTDateStr(nil))

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Length", strconv.Itoa(csv.Len()))

	_, err := io.Copy(w, csv)
	if err != nil {
		return apiError.InternalServerError(err)
	}

	return nil
}

func SendCSVResponseWithFileName(w http.ResponseWriter, fileName string, csvStr *string) *apiError.Error {
	br := bytes.NewBufferString(*csvStr)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.csv", fileName))
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Length", strconv.Itoa(len(*csvStr)))

	_, err := io.Copy(w, br)
	if err != nil {
		return apiError.InternalServerError(err)
	}

	return nil
}
