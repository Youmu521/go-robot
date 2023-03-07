package keywords

import (
	"encoding/json"
	"github.com/imroc/req/v3"
	"robot/config"
	"robot/request"
)

// ChooseMusic 点歌
func ChooseMusic(name string) {
	musicInfo := searchMusic(name)
	request.M().SendMusic(musicInfo.Id)
}

// SingMusic 唱歌  需要安装 ffmpeg
func SingMusic(name string) {
	musicInfo := searchMusic(name)
	request.M().SendRecord(musicInfo.Url)
}

func searchMusic(name string) config.List {
	var url = "https://api.acgxt.com/v1/163/music_search"
	// get 提交
	client := req.C()
	resp, err := client.R().SetQueryParams(map[string]string{
		"name": name,
		"size": "1",
	}).Get(url)
	if err != nil {
		panic("获取失败")
	}
	if !resp.IsSuccessState() {
		panic("状态码不正确")
	}
	// 转 byte
	data, err := resp.ToBytes()
	if err != nil {
		panic("转换 byte 失败")
	}
	// 转结构体
	musicSearchData := config.MusicSearchData{}
	err = json.Unmarshal(data, &musicSearchData)
	if err != nil {
		panic("转 结构体 失败")
	}
	// 判断查询是否成功
	if musicSearchData.Status == 1 {
		panic("歌曲获取失败")
	}
	return musicSearchData.Data.List[0]
}
