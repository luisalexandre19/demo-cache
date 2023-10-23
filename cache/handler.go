package cache

import (
	"crypto/sha1"
	"fmt"

	boot "demo.cache/boostrap"
	"demo.cache/cache/domain"
	"demo.cache/cache/provider/memory"
	"demo.cache/cache/provider/redis"
	log "github.com/sirupsen/logrus"
)

//Gera a key do cache no redis, gera um sha1 como key usando os parametros recebidos
func GenerateCacheKey(paramns string) string {
	var hash = sha1.New()

	hash.Write([]byte(paramns))

	hashKey := fmt.Sprintf("%x", hash.Sum(nil))

	log.Debugf("-------------------------Key  [ %s ]", hashKey)
	hash.Reset()
	return boot.APP_CONFIG.Identification + "-" + hashKey
}

func FactoryProvider() domain.CacheRepository {
	log.Infof("Cache provider : %s ", boot.APP_CONFIG.Provider)
	switch provider := boot.APP_CONFIG.Provider; provider {
	case "MEMORY":
		log.Info("Starting memory cache")
		mem := &memory.MemoryProvider{}
		return mem

	case "REDIS":
		log.Info("Starting redis cache")
		redis := &redis.RedisProvider{}
		return redis
	default:
		log.Info("Starting memory cache")
		return &memory.MemoryProvider{}
	}
}

func tryConnectionIsError() bool {
	return RepositoryCache.Error() == nil
}
