package main

import (
	"net/http"

	"demo.cache/boostrap"
	"demo.cache/cache"
	"demo.cache/web"
	log "github.com/sirupsen/logrus"
)

func main() {

	boostrap.ConfigureLogger()
	boostrap.Initialize()

	log.Infof(" ############  CACHE (DEMO)  #################")
	log.Info(" Start at port: " + boostrap.APP_CONFIG.Port)

	cache.Initialize()

	web.CreateEndpoints()

	log.Infof(" Exec on %s port", boostrap.APP_CONFIG.Port)

	log.Fatal(http.ListenAndServe(":"+boostrap.APP_CONFIG.Port, web.GetRouter))

}
