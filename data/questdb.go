package data

import (
	"log"
	"log-collector/config"
	"net"
)

var questDbConn net.Conn

func InitQuestDB() {
	if config.QuestDB == "" {
		log.Fatalln("QuestDB Missing...")
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp4", config.QuestDB)
	if err != nil {
		log.Fatalln("InitQuestDB - net.ResolveTCPAddr(\"tcp4\", config.QuestDB)():", err)
	}

	questDbConn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalln("InitQuestDB - net.DialTCP(\"tcp\", nil, tcpAddr):", err)
	}

	log.Println("InitQuestDB() - Connected")
	defer questDbConn.Close()
}

func GetQuestDBConn() *net.Conn {
	return &questDbConn
}
