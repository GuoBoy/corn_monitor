package config

import (
	"encoding/json"
	"log"
	"os"
)

// PushPlusModel pushplus配置
type PushPlusModel struct {
	Token string `json:"token"`
}

// GoldModel gold配置
type GoldModel struct {
	Threshold float64 `json:"threshold"`
}

type CfgModel struct {
	PushPlus PushPlusModel `json:"pushplus"`
	Gold     GoldModel     `json:"gold"`
}

var Cfg CfgModel

func init() {
	dt, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("读取配置文件失败")
	}
	if err = json.Unmarshal(dt, &Cfg); err != nil {
		log.Fatal("解析配置文件失败")
	}
}
