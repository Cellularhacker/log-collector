package util

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"log"
	"log-collector/service/telegram"
	"time"
)

func SendFailed(location string, err error) {
	loc, _ := time.LoadLocation("Asia/Seoul")
	at := time.Now().In(loc)
	msg := fmt.Sprintf("%v", aurora.Red(fmt.Sprintf("[ERROR/%s]\n=> %s", location, err)))

	telegram.SendMessageAt(msg, at)
}

func SendNotice(header, location, content string) {
	loc, _ := time.LoadLocation("Asia/Seoul")
	at := time.Now().In(loc)
	msg := fmt.Sprintf("[%s/%s]\n=> %s\n", header, location, content)
	log.Println(aurora.Cyan(msg))

	telegram.SendMessageAt(msg, at)
}
