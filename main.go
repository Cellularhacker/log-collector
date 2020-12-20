package main

import (
	"github.com/chyeh/pubip"
	"github.com/rs/cors"
	"log-collector/config"

	"fmt"
	"log"
	"log-collector/route"
	"log-collector/service/telegram"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var localIP = ""
var pubIP = ""
var hostname = ""

func init()  {
	telegram.Init()
}

func main()  {
	// Send Startup Message
	go func() {
		hostname, _ = os.Hostname()
		localIP = GetOutboundIP().String()
		pIP, _ := pubip.Get()
		pubIP = pIP.String()

		// Send a telegramMessage to notice server has been started.
		telegram.SendStarted(hostname, localIP, pubIP)
	}()

	// Loading HTTP Router
	router := route.Load()

	//Allowing c for now for frontend reset pass
	corsMiddle()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		MaxAge:           1728000,
	})
	handler := c.Handler(router)

	// Handling server requested to stop.
	handleServerStop()

	log.Printf("Listening on port %s\n", config.ListenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.ListenPort), handler))
}

func corsMiddle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if request.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(nil)
		}
	})
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func handleServerStop() {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		if sig == syscall.SIGTERM || sig == syscall.SIGINT {
			telegram.SendStopped(hostname, localIP, pubIP)
			os.Exit(0)
		}
	}()
}
