package server

import "mssgserver/utils"
type Ip struct {
	Id int64 `db:"id"`
	Ip string `db:"ip"`
	Status int64 `db:"status"`
	Mark string `db:"mark"`
	CreateTime string `db:"create_time"`
}

func (i *Ip) Allow (ip string) (ipinfo Ip,err error) {
	sqlStr1 := "select id, ip, status, mark, create_time from ct_white_ip where ip = ?"
	err = utils.DB.Get(&ipinfo,sqlStr1,ip)
	return
}