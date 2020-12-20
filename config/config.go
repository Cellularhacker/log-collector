package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ModeProduction  = "Production"
	ModeDevelopment = "Development"

	ServerName = "log-collector"
)

var ListenPort = ""

var SqlURL = ""
var MongoURL = ""
var MongoAddr []string
var MongoUsername = ""
var MongoAuthDB = ""
var MongoPass = ""
var MongoReplicaSetName = ""
var QuestDB = ""

var EncryptionSecret = ""

var Mode = ModeDevelopment

var Loc *time.Location

func init() {
	EncryptionSecret = os.Getenv("LOGCOLLECTOR__ENCRYPT")

	ListenPort = os.Getenv("LOGCOLLECTOR_PORT")
	if _, err := strconv.Atoi(ListenPort); err != nil {
		ListenPort = "5000"
	}

	Mode = os.Getenv("LOGCOLLECTOR__MODE")
	if IsProductionMode() {
		log.Println("Running LOGCOLLECTOR_ in Production Mode")
	} else {
		log.Println("Running LOGCOLLECTOR_ in Development Mode")
	}

	SqlURL = os.Getenv("LOGCOLLECTOR__MYSQL_URL")

	MongoURL = os.Getenv("LOGCOLLECTOR__MONGO_URL")
	MongoAddr = strings.Split(os.Getenv("LOGCOLLECTOR__MONGO_ADDR"), ",")
	MongoUsername = os.Getenv("LOGCOLLECTOR__MONGO_USERNAME")
	MongoPass = os.Getenv("LOGCOLLECTOR__MONGO_PASS")
	MongoAuthDB = os.Getenv("LOGCOLLECTOR__MONGO_AUTH_DB")
	MongoReplicaSetName = os.Getenv("LOGCOLLECTOR__MONGO_REPLICA_SET_NAME")

	QuestDB = os.Getenv("LOGCOLLECTOR_QUESTDB")

	Loc, _ = time.LoadLocation("Asia/Seoul")
}

func IsProductionMode() bool {
	return Mode == ModeProduction
}
