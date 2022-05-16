package model

import "time"

type Operator struct {
	Id         int64  `json:"Id" form:"Id"`
	Username   string `json:"Username" form:"Username"`
	Password   string `json:"Password" form:"Password"`
	Nama       string `json:"Nama" from:"Nama"`
	Created_at time.Time
	Updated_at time.Time
}
