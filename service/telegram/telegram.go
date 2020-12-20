package telegram

import (
	"fmt"
	"log"
	"log-collector/config"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var Token = ""
var MonitorChatID = ""
var BlackCowChatID = ""

func init() {
	Token = os.Getenv("PARSER_TELEGRAM_ACCESS_TOKEN")
	if Token == "" {
		log.Fatalln("PARSER_TELEGRAM_ACCESS_TOKEN missing")
	}
	MonitorChatID = os.Getenv("PARSER_TELEGRAM_CHAT_ID_MONITOR")
	if MonitorChatID == "" {
		log.Fatalln("PARSER_TELEGRAM_CHAT_ID_MONITOR missing")
	}
	BlackCowChatID = os.Getenv("PARSER_TELEGRAM_CHAT_ID_NOTICE")
	if BlackCowChatID == "" {
		log.Fatalln("PARSER_TELEGRAM_CHAT_ID_NOTICE missing")
	}
}

var to *MonitorRoom

var initialized = false

type MonitorRoom struct{}

func (*MonitorRoom) Recipient() string {
	return MonitorChatID
}

var bot *tb.Bot

func Init() {
	log.Println("Initializing telegram bot..")
	var err error
	bot, err = tb.NewBot(tb.Settings{
		Token:  Token,
		Poller: &tb.LongPoller{Timeout: 5 * time.Second},
	})
	if err != nil {
		log.Fatalln(err)
	}

	to = &MonitorRoom{}
	initialized = true
}

func SendMessage(message string) {
	loc, _ := time.LoadLocation("Asia/Seoul")

	SendMessageAt(message, time.Now().In(loc))
}

func SendMessageAt(message string,  t time.Time) {
	SendMessageAtOn(message, to, t)
}

func SendMessageAtOn(msg string, r tb.Recipient, t time.Time) {
	if !config.IsProductionMode() || !initialized {
		return
	}

	text := fmt.Sprintf("%s\n%s", msg, t.Format(time.RFC822))
	_, err := bot.Send(r, text)
	if err != nil {
		log.Printf("[Telegram](Send) Failed to send: %s\n", msg)
	}
	log.Printf("[Telegram] sent message (%s)\n", r.Recipient())
}

func SendStarted(hostname string, localIP string, pubIP string) {
	log.Println("SendStarted()")
	msg := fmt.Sprintf("<%s> started successfully\nHostname:%s\nLocal IP:%s\nPublic IP:%s\n", config.ServerName, hostname, localIP, pubIP)
	SendMessage(msg)
}

func SendStopped(hostname string, localIP, pubIP string) {
	msg := fmt.Sprintf("<%s> stopping normally\nHostname:%s\nLocal IP:%s\nPublic IP:%s", config.ServerName, hostname, localIP, pubIP)
	SendMessage(msg)
}

func SendFailed(location string, err error, at time.Time) {
	msg := fmt.Sprintf("[ERROR/%s]\n=> %s", location, err)
	SendMessageAt(msg, at)
}

func SendFailedMsg(message string, at time.Time) {
	SendMessageAt(message, at)
}
