package vo

import "time"

type HdcLog struct {
	Id         int64  `json:"id"`
	ClientHost string `json:"client_host"`
	ClientPid  string `json:"client_pid"`
}

type Student struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name" valid:"required"`
	Ago        int32     `json:"ago" valid:"required"`
	Sex        int32     `json:"sex" valid:"required"`
	CreateTime time.Time `json:"create_time"`
}
