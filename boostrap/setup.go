package boostrap

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var (
	APP_CONFIG   AppConfig
	REDIS_CONFIG RedisConfig
)

func Initialize() {

	APP_CONFIG = LoadEnvApp()
	REDIS_CONFIG = LoadEnvRedis()

}

func LoadEnvApp() AppConfig {
	var app AppConfig

	app.Port = loadEnvString("SDC_SERVER_PORT", "9200")

	app.BussinesContainerAddr = loadEnvString("SDC_BUSSINES_URL", "http://localhost:8080")

	app.Provider = loadEnvString("SDC_CACHE_PROVIDER", "MEMORY")

	app.Identification = "demo-cache"

	log.Info("Boot with confis :  ", app)
	return app

}

func LoadEnvRedis() RedisConfig {
	var redis RedisConfig
	var err error

	redis.Host = loadEnvString("SDC_REDIS_HOST", "localhost")
	redis.Port = loadEnvString("SDC_REDIS_PORT", "6379")
	redis.Password = os.Getenv("SDC_REDIS_PASSWORD")

	redis.Ttl = loadEnvInt("SDC_REDIS_TTL_MIN", 45)

	redis.User = os.Getenv("SDC_REDIS_USER")

	redis.Context = os.Getenv("SDC_REDIS_CONTEXT")

	redis.Db = loadEnvInt("SDC_REDIS_DB", 0)

	redis.PoolSize = loadEnvInt("SDC_REDIS_POOLSIZE", 15)

	redis.MaxConnAge = loadEnvInt("SDC_REDIS_MAXCONNAGE_SEC", 10)

	redis.IdleTimeout = loadEnvInt("SDC_REDIS_IDLETIMEOUT_SEC", 10)

	redis.WriteTimeout = loadEnvInt("SDC_REDIS_WRITETIMEOUT_SEC", 2)

	redis.ReadTimeout = loadEnvInt("SDC_REDIS_READTIMEOUT_SEC", 2)

	redis.DialTimeout = loadEnvInt("SDC_REDIS_DIALTIMEOUT_SEC", 1)

	redis.User = os.Getenv("SDC_REDIS_USER")

	redis.Context = os.Getenv("SDC_REDIS_CONTEXT")

	redis.Db = loadEnvInt("SDC_REDIS_DB", 0)

	redis.FailFast, err = strconv.ParseBool(loadEnvString("SDC_REDIS_FAIL_FAST", "false"))

	if err != nil { //setar "false" define  que app ira iniciar mesmo se conexao com provider falhar
		// apenas como poxy, sem cache
		redis.FailFast = false
	}

	log.Infof("Redis configs [%s] [%s] [%s] :  ", redis.Host, redis.Ttl)
	return redis
}
