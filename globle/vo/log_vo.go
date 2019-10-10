package vo

import "time"

type Log struct {
	Id         int64     `json:"id"`
	Uuid       string    `json:"uuid"`
	ClientHost string    `json:"client_host"`
	ClientPid  string    `json:"client_pid"`
	Type       int32     `json:"type"`
	Operation  string    `json:"operation"`
	Path       string    `json:"path"`
	Command    string    `json:"command"`
	Status     string    `json:"status"`
	ReturnCode int32     `json:"return_code"`
	Info       string    `json:"info"`
	ExecTime   int32     `json:"exec_time"`
	CreateTime time.Time `json:"create_time"`
}
