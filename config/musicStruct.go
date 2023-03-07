package config

type MusicSearchData struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"data"`
}

type Data struct {
	List []List `json:"list"`
	Nav  Nav    `json:"nav"`
}

type List struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Preview string `json:"preview"`
	Url     string `json:"url"`
}

type Nav struct {
	Count int `json:"count"`
	Page  int `json:"page"`
	Size  int `json:"size"`
}
