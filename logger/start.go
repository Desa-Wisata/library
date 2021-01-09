package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

// StartRecord ...
func StartRecord(req *http.Request, start time.Time) *http.Request {
	ctx := req.Context()

	v := new(Data)
	v.RequestID = uuid.New().String()
	v.Service = BaseName(filepath.Base(os.Args[0]))
	v.TimeStart = start
	v.Headers = DumpRequest(req)
	v.Method = req.Method
	v.Host = req.Host
	v.Endpoint = req.URL.Path
	v.Body = DumpRequestBody(req)
	ctx = context.WithValue(ctx, LogKey, v)

	return req.WithContext(ctx)
}

// DumpRequest is for get all data request header
func DumpRequest(req *http.Request) map[string]string {
	request := make(map[string]string)
	// Loop through headers
	for name, headers := range req.Header {
		request[name] = strings.Join(headers, ", ")
	}

	return request
}

// BaseName is for get name of file without extention
func BaseName(s string) string {
	n := strings.LastIndexByte(s, '.')
	if n == -1 {
		return s
	}

	return s[:n]
}

// DumpRequestBody is func for extract data request body
func DumpRequestBody(req *http.Request) map[string]interface{} {
	var bodyMap map[string]interface{}

	if (req.Method != http.MethodGet) && !strings.Contains(req.Header.Get("Content-Type"), "form-data") {
		reqBody, err := ioutil.ReadAll(req.Body)
		defer func(req *http.Request) {
			req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		}(req)

		if err != nil {
			return nil
		}

		err = json.Unmarshal(reqBody, &bodyMap)
		if err != nil {
			return nil
		}
	}

	return bodyMap
}
