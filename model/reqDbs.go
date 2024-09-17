package model

import "strings"

type ReqDbs struct {
	Host     string `gorm:"primaryKey" json:"host"`
	User     string `json:"user"`
	Password string `json:"pass"`
	Dbname   string `json:"dbname"`
	Port     string `json:"port"`
	Query    string `json:"query"`
}

func (reqDbs *ReqDbs) GetDbString() string {
	var s strings.Builder
	s.WriteString("host=" + reqDbs.Host)
	s.WriteString(" user=" + reqDbs.User)
	s.WriteString(" password=" + reqDbs.Password)
	s.WriteString(" dbname=" + reqDbs.Dbname)
	s.WriteString(" port=" + reqDbs.Port + " sslmode=disable TimeZone=Asia/Shanghai")
	return s.String()
}

func (reqDbs *ReqDbs) GetQuery() string {
	return reqDbs.Query
}
