package model

// 这是爬虫爬好的数据写进infoweb数据库的recruit_info数据表中

type Info struct {
	Title   string
	Salary  string
	Company string
	City    string
	Time    string
	Link    string
}

func (Info) TableName() string {
	return "recruit_info"
}
