package notice

import (
	"fmt"
	"github.com/corn_monitor/config"
	"github.com/corn_monitor/db"
	"io"
	"net/http"
	"net/url"
)

func Pushplus(title, content string) {
	log := db.LogModel{}
	api := "http://www.pushplus.plus/send"
	data := url.Values{"token": {config.Cfg.PushPlus.Token},
		"title": {title}, "content": {content}}
	resp, err := http.PostForm(api, data)
	if err != nil {
		log.Add("send", title, content, err.Error())
		return
	}
	log.Add("send", title, content, "is ok")
	res, _ := io.ReadAll(resp.Body)
	fmt.Println(string(res))
}
