package route

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log-collector/controller/api/v1"
	"log-collector/route/middleware/logRequest"
	"log-collector/util"
	"net/http"
)

func Load() http.Handler {
	return middleware(routes())
}

func routes() *httprouter.Router {
	router := httprouter.New()

	// MARK: SystemLogs
	router.GET(api1("/system_logs"), handle(v1.SystemLogsGET))
	router.POST(api1("/system_logs"), handle(v1.SystemLogsPOST))

	// MARK: Utilities & Easter Eggs.
	router.GET(api1("/coffee"), handle(v1.CoffeeGET))
	router.GET(api1("/ping"), handle(v1.PingGET))

	return router
}

//handle wraps out api handler to router
func handle(handler util.APIHandler) httprouter.Handle {
	return handler.Handle()
}

func middleware(h http.Handler) http.Handler {
	h = logRequest.Handler(h)

	return h
}

func api1(link string) string {
	return fmt.Sprintf("/v1%s", link)
}
