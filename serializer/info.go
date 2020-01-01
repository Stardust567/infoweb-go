package serializer

import "infoweb/model"

type Info struct {
	Title   string `json:"title"`
	Salary  string `json:"salary"`
	Company string `json:"company"`
	City    string `json:"city"`
	Time    string `json:"time"`
	Link    string `json:"link"`
}

func BuildInfo(info model.Info) Info {
	return Info{
		Title:   info.Title,
		Salary:  info.Salary,
		Company: info.Company,
		City:    info.City,
		Time:    info.Time,
		Link:    info.Link,
	}
}

// BuildUserResponse 序列化用户响应
func BuildInfoResponse(info model.Info) Response {
	return Response{
		Data: BuildInfo(info),
	}
}

// BuildVideos 序列化视频列表
func BuildInfoList(items []model.Info) (InfoList []Info) {
	for _, item := range items {
		info := BuildInfo(item)
		InfoList = append(InfoList, info)
	}
	return InfoList
}
