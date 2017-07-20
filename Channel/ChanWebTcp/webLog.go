package ChanWebTcp

import (
	"fmt"
	"sku/WebServer/WebKey"
	"time"
)

type LogObj struct {
	AtTime  string      `json:"time"`
	LogType string      `json:"type"`
	Content interface{} `json:"content"`
}

func SendWebLog(logType string, info interface{}) {
	log := new(LogObj)

	log.AtTime = fmt.Sprintf("%d:%d:%d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	log.LogType = logType
	log.Content = info

	SendWeb(WebKey.WEB_CLIENT_LOG, log)
}
