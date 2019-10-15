package vo

import "time"

type EtcdRegisterVo struct {
	Endpoint   string    `json:"endpoint"`
	UpdateTime time.Time `json:"update_time"`
}
