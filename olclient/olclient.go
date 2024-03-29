package olclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func GetLastError() string {
	projectId := os.Getenv("OPENLOG_PROJECT_ID")
	uri := fmt.Sprintf("/openlog/api/v1/logs?size=1&severity=error&projectId=%s&orderBy=date", projectId)

	logResponse, err := httpRequest[LogResponse](uri, http.MethodGet, nil)
	if err != nil {
		return "Something went wrong during api request"
	}

	if logResponse != nil && len(logResponse.Logs) > 0 {
		log := logResponse.Logs[0]
		return fmt.Sprintf("LAST ERROR:\n\nProject Id: %v\nHostname: %s\nDate: %s\nSeverity: %s\nCode: %s\nAction: %s\nMessage: %s",
			log.ProjectId,
			log.Hostname,
			log.Date,
			log.Severity,
			log.Code,
			log.Action,
			log.Message)
	}
	return "No errors were found"
}

func httpRequest[T OpenlogResponses](uri string, method string, body io.Reader) (*T, error) {
	const CONNECTION_TIMEOUT = 10
	client := &http.Client{Timeout: CONNECTION_TIMEOUT * time.Second}

	endpoint := fmt.Sprintf("%s%s", os.Getenv("OPENLOG_API_BASE_PATH"), uri)
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		log.Print("Could not create request", err)
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Print("Could not read response body", err)
		return nil, err
	}
	defer res.Body.Close()

	data := new(T)
	json.NewDecoder(res.Body).Decode(data)
	return data, nil
}
