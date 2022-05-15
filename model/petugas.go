package model

import "time"

type Operator struct {
	Id         int64  `json:"Id" form:"Id"`
	Username   string `json:"username" form:"username"`
	Password   string `json:"password" form:"password"`
	Nama       string `json:"Nama" from:"Nama"`
	Created_at time.Time
	Updated_at time.Time
}
