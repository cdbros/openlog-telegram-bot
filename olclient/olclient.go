package olclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Log struct {
	ProjectId int
	Hostname  string
	Date      string
	Severity  string
	Code      string
	Action    string
	Message   string
}

type LogResponse struct {
	TotalPages    int
	CurrentPage   int
	TotalElements int
	Size          int
	Logs          []Log
}

type OpenlogResponses interface {
	LogResponse
}

const CONNECTION_TIMEOUT = 10

func GetLastError() string {
	var projectId = os.Getenv("OPENLOG_PROJECT_ID")
	var uri = fmt.Sprintf("/openlog/api/v1/logs?size=1&severity=error&projectId=%s&orderBy=date", projectId)
	var logResponse = httpRequest(uri, http.MethodGet, nil)

	if logResponse != nil && len(logResponse.Logs) > 0 {
		var log = logResponse.Logs[0]
		return fmt.Sprintf("LAST ERROR:\n\nProject Id: %s\nHostname: %s\nDate: %s\nSeverity: %s\nCode: %s\nAction: %s\nMessage: %s",
			fmt.Sprint(log.ProjectId),
			log.Hostname,
			log.Date,
			log.Severity,
			log.Code,
			log.Action,
			log.Message)
	}
	return "No errors were found"
}

func httpRequest[T OpenlogResponses](uri string, method string, body io.Reader) *T {
	var client = &http.Client{Timeout: CONNECTION_TIMEOUT * time.Second}

	var endpoint = fmt.Sprintf("%s%s", os.Getenv("OPENLOG_API_BASE_PATH"), uri)
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		fmt.Printf("Could not create request: %s\n", err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Could not read response body: %s\n", err)
	}
	defer res.Body.Close()

	var data = new(T)
	json.NewDecoder(res.Body).Decode(data)
	return data
}
