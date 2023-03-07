package gold

import (
	"github.com/corn_monitor/config"
	"github.com/corn_monitor/db"
	"github.com/corn_monitor/notice"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

// Run 获取每日金价并存储
func Run() {
	logS := db.LogModel{}
	api := "http://www.dyhjw.com/hjtd"
	req, _ := http.NewRequest(http.MethodGet, api, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.63")
	for i := 0; i < 5; i++ {
		logS.Add("this is the", i, "times request")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			logS.Add("request err", err.Error())
			continue
		}
		f, _ := os.Create("temp.html")
		temp, _ := io.ReadAll(resp.Body)
		f.Write(temp)
		f.Close()
		regx := `今开：<font>([\d.]+)<`
		tempRes := regexp.MustCompile(regx).FindStringSubmatch(string(temp))
		if len(tempRes) != 0 {
			res, err := strconv.ParseFloat(tempRes[1], 64)
			if err != nil {
				logS.Add("parse res err", err.Error())
			}
			if res <= config.Cfg.Gold.Threshold {
				// 推送
				notice.Pushplus("金价提醒", "price is "+tempRes[1])
				logS.Add("push message, the value is", res)
				return
			}
			logS.Add("today value is", res)
			return
		} else {
			logS.Add("the res is empty")
		}
	}
}
