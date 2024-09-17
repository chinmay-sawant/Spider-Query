package model

import "strings"

type Dbs struct {
	Host     string `gorm:"primaryKey" json:"host"`
	User     string `json:"user"`
	Password string `json:"pass"`
	Dbname   string `json:"dbname"`
	Port     string `json:"port"`
}

func (dbs *Dbs) GetDbString() string {
	var s strings.Builder
	s.WriteString("host=" + dbs.Host)
	s.WriteString(" user=" + dbs.User)
	s.WriteString(" password=" + dbs.Password)
	s.WriteString(" dbname=" + dbs.Dbname)
	s.WriteString(" port=" + dbs.Port + " sslmode=disable TimeZone=Asia/Shanghai")
	return s.String()
}

// "host=localhost user=sc password=sc dbname=sc port=5432 sslmode=disable TimeZone=Asia/Shanghai"
