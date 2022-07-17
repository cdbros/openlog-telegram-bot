package olclient

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
