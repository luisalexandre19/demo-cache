package boostrap

import (
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ConfigureLogger() {

	logLevel, err := log.ParseLevel(loadEnvString("SDC_LOG_LEVEL", "info"))
	if err != nil {
		log.Errorf("Error on load configs of log %s, using defaul - info", err)
		logLevel = log.ErrorLevel
	}

	log.SetLevel(logLevel)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			pathArray := strings.Split(f.File, "/")
			file = " - " + pathArray[len(pathArray)-1]

			funcArray := strings.Split(f.Function, "/")
			function = " : " + funcArray[len(funcArray)-1] + " : "

			return
		},
		PadLevelText: true,
	})

}
